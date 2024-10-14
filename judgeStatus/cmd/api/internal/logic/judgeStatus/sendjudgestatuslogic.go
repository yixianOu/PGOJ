package judgeStatus

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"oj-micro/common/dataType"
	"oj-micro/common/xcode"
	"oj-micro/judgeStatus/cmd/api/internal/svc"
	"oj-micro/judgeStatus/cmd/api/internal/types"
	"oj-micro/judgeStatus/cmd/rpc/pb"
	problemPb "oj-micro/problems/cmd/rpc/pb"
	"strconv"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendJudgeStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSendJudgeStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendJudgeStatusLogic {
	return &SendJudgeStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SendJudgeStatusLogic) SendJudgeStatus(req *types.SendJudgeStatusRequest) (resp *types.SendJudgeStatusResponse, err error) {
	expireTime, err := l.ctx.Value("access_expire").(json.Number).Int64()
	if err != nil || time.Now().Unix() > expireTime {
		return nil, xcode.UnauthorizedTokenTimeout
	}
	userID, err := l.ctx.Value("user_id").(json.Number).Int64()
	if err != nil {
		return nil, xcode.UnauthorizedUserNotLogin
	}

	//查找题目
	problemById, err := l.svcCtx.ProblemServiceRpc.GetProblemById(l.ctx, &problemPb.GetProblemByIdReq{
		Id: req.ProblemId,
	})
	if err != nil {
		return nil, err
	}
	//查找测试用例
	testcases, err := l.svcCtx.ProblemServiceRpc.SearchTestcases(l.ctx, &problemPb.SearchTestcasesReq{
		ProblemId: req.ProblemId,
	})
	if err != nil {
		return nil, err
	}

	//向评测机发送评测请求
	stream, err := l.svcCtx.JudgeServiceRpc.AddJudgestatus(l.ctx, &pb.AddJudgestatusReq{
		UserId:         userID,
		ProblemId:      req.ProblemId,
		ProblemTitle:   req.ProblemTitle,
		Oj:             req.Oj,
		Language:       req.Language,
		Code:           req.Code,
		Length:         req.Length,
		Contest:        req.Contest,
		ContestProblem: req.ContestProblem,
		Rating:         req.Rating,
		Ip:             req.Ip,

		ProblemCode: problemById.Problem.ProblemCode,
		TimeLimit:   problemById.Problem.Time,
		MemoryLimit: problemById.Problem.Memory,

		CaseNum: int64(len(testcases.Testcases)),
	})
	if err != nil {
		return nil, err
	}

	//创建评测记录
	updateJudgestatusReq := &pb.UpdateJudgestatusReq{
		Result:       string(dataType.Accept),
		TimeCost:     0,
		MemoryCost:   0,
		Message:      "",
		InputData:    "",
		SampleOutput: "",
		UserOutput:   "",
		TestCases:    "",
	}
	//处理评测结果
	var result *pb.AddJudgestatusResp
	for {
		result, err = stream.Recv()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			} else {
				logx.Errorf("stream.Recv error: %v", err)
				return nil, xcode.ServerErr
			}
		}

		if result.Result != string(dataType.Accept) {
			if updateJudgestatusReq.Result == string(dataType.Accept) {
				updateJudgestatusReq.Result = result.Result
				updateJudgestatusReq.Message = result.Message
				updateJudgestatusReq.InputData = result.InputData
				updateJudgestatusReq.SampleOutput = result.SampleOutput
				updateJudgestatusReq.UserOutput = result.UserOutput
			}
			updateJudgestatusReq.TestCases = appendStringsWithSemicolon(updateJudgestatusReq.TestCases, strconv.FormatInt(result.TestCase, 10))
		}
		updateJudgestatusReq.TimeCost = max(updateJudgestatusReq.TimeCost, result.TimeCost)
		updateJudgestatusReq.MemoryCost = max(updateJudgestatusReq.MemoryCost, result.MemoryCost)
	}
	updateJudgestatusReq.JudgeId = result.JudgeId
	//更新评测记录
	_, err = l.svcCtx.JudgeServiceRpc.UpdateJudgestatus(l.ctx, updateJudgestatusReq)
	if err != nil {
		return nil, err
	}

	//创建题目数据
	updateProblemdataReq := problemPb.UpdateProblemdataReq{
		ProblemdataId: req.ProblemId,
		Submission:    1,
	}
	switch updateJudgestatusReq.Result {
	case string(dataType.Accept):
		updateProblemdataReq.Ac = 1
	case string(dataType.MemoryLimitExceeded):
		updateProblemdataReq.Mle = 1
	case string(dataType.TimeLimitExceeded):
		updateProblemdataReq.Tle = 1
	case string(dataType.RuntimeError):
		updateProblemdataReq.Rte = 1
	case string(dataType.WrongAnswer):
		updateProblemdataReq.Wa = 1
	case string(dataType.OutputLimitExceeded):
		updateProblemdataReq.Ole = 1
	case string(dataType.CompileError):
		updateProblemdataReq.Ce = 1
	case string(dataType.UnknownError):
		updateProblemdataReq.Ue = 1
	case string(dataType.FloatError):
		updateProblemdataReq.Fe = 1
	case string(dataType.SegmentationFault):
		updateProblemdataReq.Sf = 1
	default:
		logx.Errorf("unknown result: %v", updateJudgestatusReq.Result)
		return nil, xcode.ServerErr
	}
	//更新题目数据
	_, err = l.svcCtx.ProblemServiceRpc.UpdateProblemdata(l.ctx, &updateProblemdataReq)
	if err != nil {
		return nil, err
	}
	//获取评测结果
	judgestatusById, err := l.svcCtx.JudgeServiceRpc.GetJudgestatusById(l.ctx, &pb.GetJudgestatusByIdReq{
		Id: updateJudgestatusReq.JudgeId,
	})
	if err != nil {
		return nil, err
	}
	//返回评测结果
	resp = &types.SendJudgeStatusResponse{
		JudgeStatus: types.JudgeStatus{
			JudgeId:        judgestatusById.Judgestatus.JudgeId,
			UserId:         judgestatusById.Judgestatus.JudgeId,
			ProblemId:      judgestatusById.Judgestatus.ProblemId,
			ProblemTitle:   judgestatusById.Judgestatus.ProblemTitle,
			Oj:             judgestatusById.Judgestatus.Oj,
			Result:         judgestatusById.Judgestatus.Result,
			TimeCost:       judgestatusById.Judgestatus.TimeCost,
			MemoryCost:     judgestatusById.Judgestatus.MemoryCost,
			Length:         judgestatusById.Judgestatus.Length,
			Language:       judgestatusById.Judgestatus.Language,
			Code:           judgestatusById.Judgestatus.Code,
			SubmitTime:     judgestatusById.Judgestatus.SubmitTime,
			Judger:         judgestatusById.Judgestatus.Judger,
			Contest:        judgestatusById.Judgestatus.Contest,
			ContestProblem: judgestatusById.Judgestatus.ContestProblem,
			TestCases:      judgestatusById.Judgestatus.TestCases,
			Message:        judgestatusById.Judgestatus.Message,
			Rating:         judgestatusById.Judgestatus.Rating,
			Ip:             judgestatusById.Judgestatus.Ip,
			InputData:      judgestatusById.Judgestatus.InputData,
			SampleOutput:   judgestatusById.Judgestatus.SampleOutput,
			UserOutput:     judgestatusById.Judgestatus.UserOutput,
		},
	}
	return
}
func appendStringsWithSemicolon(stringsToAppend ...string) string {
	// 使用 strings.Join 方法将多个字符串用 ";" 分隔拼接起来
	return strings.Join(stringsToAppend, ";")
}
