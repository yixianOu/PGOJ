package judgeStatus

import (
	"context"
	"encoding/json"
	"oj-micro/common/xcode"
	"oj-micro/judgeStatus/cmd/api/internal/svc"
	"oj-micro/judgeStatus/cmd/api/internal/types"
	"oj-micro/judgeStatus/cmd/rpc/pb"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetJudgeStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetJudgeStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetJudgeStatusLogic {
	return &GetJudgeStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetJudgeStatusLogic) GetJudgeStatus(req *types.GetJudgeStatusRequest) (resp *types.GetJudgeStatusResponse, err error) {
	expireTime, err := l.ctx.Value("access_expire").(json.Number).Int64()
	if err != nil || time.Now().Unix() > expireTime {
		return nil, xcode.UnauthorizedTokenTimeout
	}

	result, err := l.svcCtx.JudgeServiceRpc.GetJudgestatusById(l.ctx, &pb.GetJudgestatusByIdReq{
		Id: req.JudgeId,
	})
	if err != nil {
		return nil, err
	}

	userID, err := l.ctx.Value("user_id").(json.Number).Int64()
	if err != nil || userID != result.Judgestatus.UserId {
		return nil, xcode.UnauthorizedUserNotLogin
	}

	resp = &types.GetJudgeStatusResponse{
		JudgeStatus: types.JudgeStatus{
			JudgeId:        result.Judgestatus.JudgeId,
			UserId:         result.Judgestatus.UserId,
			ProblemId:      result.Judgestatus.ProblemId,
			ProblemTitle:   result.Judgestatus.ProblemTitle,
			Oj:             result.Judgestatus.Oj,
			Result:         result.Judgestatus.Result,
			TimeCost:       result.Judgestatus.TimeCost,
			MemoryCost:     result.Judgestatus.MemoryCost,
			Length:         result.Judgestatus.Length,
			Language:       result.Judgestatus.Language,
			SubmitTime:     result.Judgestatus.SubmitTime,
			Code:           result.Judgestatus.Code,
			Judger:         result.Judgestatus.Judger,
			Contest:        result.Judgestatus.Contest,
			ContestProblem: result.Judgestatus.ContestProblem,
			TestCases:      result.Judgestatus.TestCases,
			Message:        result.Judgestatus.Message,
			Rating:         result.Judgestatus.Rating,
			Ip:             result.Judgestatus.Ip,
			InputData:      result.Judgestatus.InputData,
			SampleOutput:   result.Judgestatus.SampleOutput,
			UserOutput:     result.Judgestatus.UserOutput,
		},
	}
	return
}
