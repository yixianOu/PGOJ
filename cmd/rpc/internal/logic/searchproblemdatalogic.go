package logic

import (
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"oj-micro/common/xcode"
	"oj-micro/problems/cmd/rpc/internal/svc"
	"oj-micro/problems/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchProblemdataLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchProblemdataLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchProblemdataLogic {
	return &SearchProblemdataLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchProblemdataLogic) SearchProblemdata(in *pb.SearchProblemdataReq) (*pb.SearchProblemdataResp, error) {
	builder := l.svcCtx.ProblemdataModel.SelectBuilder()
	result, err := l.svcCtx.ProblemdataModel.SearchProblemdataByFields(l.ctx, builder, in.Page, in.Limit, in.ScoreFloor, in.ScoreCeil, in.Auth, in.Order)
	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, xcode.RecordNotFound
		}
		logx.Errorf("search problemdata by fields fail, err : %v, result : %+v", err, result)
		return nil, xcode.ServerErr
	}

	problemdata := make([]*pb.Problemdata, 0)
	for _, v := range result {
		problemdata = append(problemdata, &pb.Problemdata{
			ProblemdataId: v.ProblemdataId,
			ProblemId:     v.ProblemId,
			Score:         v.Score,
			Auth:          v.Auth,
			Submission:    v.Submission,
			Ac:            v.Ac,
			Mle:           v.Mle,
			Tle:           v.Tle,
			Ole:           v.Ole,
			Ce:            v.Ce,
			Ue:            v.Ue,
			Rte:           v.Rte,
			Wa:            v.Wa,
		})
	}
	return &pb.SearchProblemdataResp{Problemdata: problemdata}, nil
}
