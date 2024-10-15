package logic

import (
	"context"
	"oj-micro/common/xcode"
	"oj-micro/problems/model"

	"oj-micro/problems/cmd/rpc/internal/svc"
	"oj-micro/problems/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateTestcasesLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateTestcasesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateTestcasesLogic {
	return &UpdateTestcasesLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateTestcasesLogic) UpdateTestcases(in *pb.UpdateTestcasesReq) (*pb.UpdateTestcasesResp, error) {
	err := l.svcCtx.TestCasesModel.PartialUpdate(l.ctx, &model.Testcases{
		TestId:         in.TestId,
		InputFileName:  in.InputFileName,
		OutputFileName: in.OutputFileName,
	})
	if err != nil {
		logx.Errorf("update testcases fail, err : %v", err)
		return nil, xcode.ServerErr
	}
	return &pb.UpdateTestcasesResp{}, nil
}
