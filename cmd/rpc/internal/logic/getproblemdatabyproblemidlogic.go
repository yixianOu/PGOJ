package logic

import (
	"context"
	"errors"
	"oj-micro/common/xcode"
	"oj-micro/problems/model"

	"oj-micro/problems/cmd/rpc/internal/svc"
	"oj-micro/problems/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetProblemdataByProblemIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetProblemdataByProblemIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetProblemdataByProblemIdLogic {
	return &GetProblemdataByProblemIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetProblemdataByProblemIdLogic) GetProblemdataByProblemId(in *pb.GetProblemdataByProblemIdReq) (*pb.GetProblemdataByProblemIdResp, error) {
	oneByProblemId, err := l.svcCtx.ProblemdataModel.FindOneByProblemId(l.ctx, in.ProblemId)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return nil, xcode.RecordNotFound
		}
		l.Logger.Errorf("find problemdata by problemId fail, err : %v, result : %+v", err, oneByProblemId)
		return nil, xcode.ServerErr
	}

	return &pb.GetProblemdataByProblemIdResp{Problemdata: &pb.Problemdata{
		ProblemdataId: oneByProblemId.ProblemdataId,
		ProblemId:     oneByProblemId.ProblemId,
		Score:         oneByProblemId.Score,
		Auth:          oneByProblemId.Auth,
		Submission:    oneByProblemId.Submission,
		Ac:            oneByProblemId.Ac,
		Mle:           oneByProblemId.Mle,
		Tle:           oneByProblemId.Tle,
		Ole:           oneByProblemId.Ole,
		Ce:            oneByProblemId.Ce,
		Ue:            oneByProblemId.Ue,
		Rte:           oneByProblemId.Rte,
		Wa:            oneByProblemId.Wa,
		Fe:            oneByProblemId.Fe,
		Sf:            oneByProblemId.Sf,
	}}, nil
}
