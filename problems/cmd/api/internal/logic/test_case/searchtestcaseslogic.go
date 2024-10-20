package test_case

import (
	"context"
	"oj-micro/problems/cmd/rpc/pb"

	"oj-micro/problems/cmd/api/internal/svc"
	"oj-micro/problems/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchTestCasesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSearchTestCasesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchTestCasesLogic {
	return &SearchTestCasesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SearchTestCasesLogic) SearchTestCases(req *types.SearchTestCaseRequest) (resp *types.SearchTestCaseResponse, err error) {
	testcases, err := l.svcCtx.ProblemServiceRpc.SearchTestcases(l.ctx, &pb.SearchTestcasesReq{
		ProblemId: req.ProblemId,
		TestGroup: req.TestGroup,
	})
	if err != nil {
		return nil, err
	}

	var testCases []types.TestCases
	for _, testcase := range testcases.Testcases {
		testCases = append(testCases, types.TestCases{
			TestId:             testcase.TestId,
			ProblemId:          testcase.ProblemId,
			TestGroup:          testcase.TestGroup,
			TestInputFileName:  testcase.InputFileName,
			TestOutputFileName: testcase.OutputFileName,
			UpdateAt:           testcase.UpdateAt.AsTime().Unix(),
		})
	}

	resp = &types.SearchTestCaseResponse{
		TestCases: testCases,
	}
	return
}
