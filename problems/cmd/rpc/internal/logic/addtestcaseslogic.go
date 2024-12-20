package logic

import (
	"context"
	"errors"
	"oj-micro/common/xcode"
	"oj-micro/problems/cmd/rpc/internal/code"
	"oj-micro/problems/model"

	"oj-micro/problems/cmd/rpc/internal/svc"
	"oj-micro/problems/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddTestcasesLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddTestcasesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddTestcasesLogic {
	return &AddTestcasesLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddTestcasesLogic) AddTestcases(in *pb.AddTestcasesReq) (*pb.AddTestcasesResp, error) {
	_, err := l.svcCtx.TestCasesModel.FindOneByProblemIdTestGroup(l.ctx, in.ProblemId, in.TestGroup)
	if !errors.Is(err, model.ErrNotFound) {
		return nil, code.ProblemSampleExist
	}

	result, err := l.svcCtx.TestCasesModel.Insert(l.ctx, &model.Testcases{
		ProblemId:      in.ProblemId,
		TestGroup:      in.TestGroup,
		InputFilePath:  in.InputFileName,
		OutputFilePath: in.OutputFileName,
	})
	if err != nil {
		l.Logger.Errorf("insert testcases fail, err : %v, result : %+v", err, result)
		return nil, xcode.ServerErr
	}
	return &pb.AddTestcasesResp{}, nil
}
