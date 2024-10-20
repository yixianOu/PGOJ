package test_case

import (
	"context"
	"oj-micro/problems/cmd/rpc/pb"

	"oj-micro/problems/cmd/api/internal/svc"
	"oj-micro/problems/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTestCaseByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetTestCaseByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTestCaseByIdLogic {
	return &GetTestCaseByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetTestCaseByIdLogic) GetTestCaseById(req *types.GetTestCaseByIdRequest) (resp *types.GetTestCaseByIdResponse, err error) {
	testcasesById, err := l.svcCtx.ProblemServiceRpc.GetTestcasesById(l.ctx, &pb.GetTestcasesByIdReq{
		Id: req.TestId,
	})
	if err != nil {
		return nil, err
	}

	resp = &types.GetTestCaseByIdResponse{
		TestCase: types.TestCases{
			TestId:             testcasesById.Testcases.TestId,
			ProblemId:          testcasesById.Testcases.ProblemId,
			TestGroup:          testcasesById.Testcases.TestGroup,
			TestInputFileName:  testcasesById.Testcases.InputFileName,
			TestOutputFileName: testcasesById.Testcases.OutputFileName,
			UpdateAt:           testcasesById.Testcases.UpdateAt.AsTime().Unix(),
		},
	}
	return
}
