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

type UpdateProblemdataLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateProblemdataLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateProblemdataLogic {
	return &UpdateProblemdataLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateProblemdataLogic) UpdateProblemdata(in *pb.UpdateProblemdataReq) (*pb.UpdateProblemdataResp, error) {
	newData := &model.Problemdata{
		ProblemdataId: in.ProblemdataId,
		Auth:          in.Auth,
		Score:         in.Score,
		Submission:    in.Submission,
		Ac:            in.Ac,
		Mle:           in.Mle,
		Tle:           in.Tle,
		Wa:            in.Wa,
		Rte:           in.Rte,
		Ce:            in.Ce,
		Ole:           in.Ole,
		Ue:            in.Ue,
	}

	err := l.svcCtx.ProblemdataModel.PartialUpdate(l.ctx, newData)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return nil, xcode.RecordNotFound
		}
		l.Logger.Errorf("update problemdata fail, err : %v, newData : %+v", err, newData)
		return nil, xcode.ServerErr
	}
	return &pb.UpdateProblemdataResp{}, nil
}
