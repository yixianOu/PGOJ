package logic

import (
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"oj-micro/common/xcode"
	"oj-micro/problems/cmd/rpc/internal/code"

	"oj-micro/problems/cmd/rpc/internal/svc"
	"oj-micro/problems/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelTestcasesLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelTestcasesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelTestcasesLogic {
	return &DelTestcasesLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelTestcasesLogic) DelTestcases(in *pb.DelTestcasesReq) (*pb.DelTestcasesResp, error) {
	err := l.svcCtx.TestCasesModel.Delete(l.ctx, in.Id)
	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, code.ProblemSampleNotExist
		}
		logx.Errorf("delete testcases fail, err : %v", err)
		return nil, xcode.ServerErr
	}
	return &pb.DelTestcasesResp{}, nil
}
