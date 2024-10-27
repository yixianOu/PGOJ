package logic

import (
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"oj-micro/common/xcode"
	"oj-micro/judgeStatus/cmd/rpc/internal/svc"
	"oj-micro/judgeStatus/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchJudgestatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchJudgestatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchJudgestatusLogic {
	return &SearchJudgestatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchJudgestatusLogic) SearchJudgestatus(in *pb.SearchJudgestatusReq) (*pb.SearchJudgestatusResp, error) {
	builder := l.svcCtx.JudgeStatusModel.SelectBuilder()
	results, err := l.svcCtx.JudgeStatusModel.SearchJudgestatusByFields(l.ctx, builder, in.Page, in.Limit, in.UserId, in.ProblemId, in.Result, in.Language, in.SubmitTime, in.Contest, in.ProblemTitle, in.Order)
	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, xcode.RecordNotFound
		}
		logx.Errorf("JugeStatus FindOne error: %v", err)
		return nil, xcode.ServerErr
	}

	var Judgestatus []*pb.Judgestatus
	for _, result := range results {
		Judgestatus = append(Judgestatus, &pb.Judgestatus{
			JudgeId:        result.JudgeId,
			UserId:         result.UserId,
			ProblemId:      result.ProblemId,
			Oj:             result.Oj,
			Result:         result.Result,
			TimeCost:       result.Time,
			MemoryCost:     result.Memory,
			Length:         result.Length,
			Language:       result.Language,
			Code:           result.Code,
			SubmitTime:     result.CreateTime.Unix(),
			Judger:         result.Judger,
			Contest:        result.Contest,
			ContestProblem: result.Contestproblem,
			TestCases:      result.Testcase,
			Ip:             result.Ip,
			Message:        result.Message,
			ProblemTitle:   result.Problemtitle,
			Rating:         result.Rating,
			InputData:      result.InputData.String,
			SampleOutput:   result.SampleOutPut.String,
			UserOutput:     result.UserOutPut.String,
		})
	}
	return &pb.SearchJudgestatusResp{Judgestatus: Judgestatus}, nil
}
