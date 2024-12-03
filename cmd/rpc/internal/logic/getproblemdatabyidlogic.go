package logic

import (
	"context"
	"errors"
	"oj-micro/common/xcode"
	"oj-micro/problems/cmd/rpc/internal/svc"
	"oj-micro/problems/cmd/rpc/pb"
	"oj-micro/problems/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetProblemdataByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetProblemdataByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetProblemdataByIdLogic {
	return &GetProblemdataByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetProblemdataByIdLogic) GetProblemdataById(in *pb.GetProblemdataByIdReq) (*pb.GetProblemdataByIdResp, error) {
	result, err := l.svcCtx.ProblemdataModel.FindOne(l.ctx, in.Id)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return nil, xcode.RecordNotFound
		}
		l.Logger.Errorf("find problemdata by id fail, err : %v, result : %+v", err, result)
		return nil, xcode.ServerErr
	}

	return &pb.GetProblemdataByIdResp{Problemdata: &pb.Problemdata{
		ProblemdataId: result.ProblemdataId,
		ProblemId:     result.ProblemId,
		Score:         result.Score,
		Auth:          result.Auth,
		Submission:    result.Submission,
		Ac:            result.Ac,
		Mle:           result.Mle,
		Tle:           result.Tle,
		Ole:           result.Ole,
		Ce:            result.Ce,
		Ue:            result.Ue,
		Rte:           result.Rte,
		Wa:            result.Wa,
		Sf:            result.Sf,
		Fe:            result.Fe,
	}}, nil
}
