package logic

import (
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"google.golang.org/protobuf/types/known/timestamppb"
	"oj-micro/common/xcode"

	"oj-micro/problems/cmd/rpc/internal/svc"
	"oj-micro/problems/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTestcasesByProblemIdAndTestGroupLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetTestcasesByProblemIdAndTestGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTestcasesByProblemIdAndTestGroupLogic {
	return &GetTestcasesByProblemIdAndTestGroupLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetTestcasesByProblemIdAndTestGroupLogic) GetTestcasesByProblemIdAndTestGroup(in *pb.GetTestcasesByProblemIdAndTestGroupReq) (*pb.GetTestcasesByProblemIdAndTestGroupResp, error) {
	testId, err := l.svcCtx.TestCasesModel.FindOneByProblemIdTestGroup(l.ctx, in.ProblemId, in.TestGroup)
	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, xcode.RecordNotFound
		}
		logx.Errorf("find testcases fail, err : %v, result : %+v", err, testId)
		return nil, xcode.ServerErr
	}

	return &pb.GetTestcasesByProblemIdAndTestGroupResp{
		Testcases: &pb.Testcases{
			TestId:         testId.TestId,
			ProblemId:      testId.ProblemId,
			TestGroup:      testId.TestGroup,
			InputFileName:  testId.InputFilePath,
			OutputFileName: testId.OutputFilePath,
			UpdateAt:       &timestamppb.Timestamp{Seconds: testId.UpdateTime.Unix()},
		},
	}, nil
}
