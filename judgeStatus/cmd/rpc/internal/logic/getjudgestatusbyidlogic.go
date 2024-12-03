package logic

import (
	"context"
	"errors"
	"oj-micro/common/xcode"
	"oj-micro/judgeStatus/cmd/rpc/internal/svc"
	"oj-micro/judgeStatus/cmd/rpc/pb"
	"oj-micro/judgeStatus/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetJudgestatusByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetJudgestatusByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetJudgestatusByIdLogic {
	return &GetJudgestatusByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetJudgestatusByIdLogic) GetJudgestatusById(in *pb.GetJudgestatusByIdReq) (*pb.GetJudgestatusByIdResp, error) {
	result, err := l.svcCtx.JudgeStatusModel.FindOne(l.ctx, in.Id)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return nil, xcode.RecordNotFound
		}
		l.Logger.Errorf("JugeStatus FindOne error: %v", err)
		return nil, xcode.ServerErr
	}

	return &pb.GetJudgestatusByIdResp{Judgestatus: &pb.Judgestatus{
		JudgeId:        result.JudgeId,
		UserId:         result.UserId,
		ProblemId:      result.ProblemId,
		Oj:             result.Oj,
		Result:         result.Result,
		TimeCost:       result.Time,
		MemoryCost:     result.Memory,
		Length:         result.Length,
		Language:       result.Language,
		SubmitTime:     result.CreateTime.Unix(),
		Judger:         result.Judger,
		Contest:        result.Contest,
		ContestProblem: result.Contestproblem,
		Code:           result.Code,
		TestCases:      result.Testcase,
		Message:        result.Message,
		ProblemTitle:   result.Problemtitle,
		Rating:         result.Rating,
		Ip:             result.Ip,
		InputData:      result.InputData.String,
		SampleOutput:   result.SampleOutPut.String,
		UserOutput:     result.UserOutPut.String,
	}}, nil
}
