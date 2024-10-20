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

type SearchTestcasesLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchTestcasesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchTestcasesLogic {
	return &SearchTestcasesLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchTestcasesLogic) SearchTestcases(in *pb.SearchTestcasesReq) (*pb.SearchTestcasesResp, error) {
	builder := l.svcCtx.TestCasesModel.SelectBuilder()
	result, err := l.svcCtx.TestCasesModel.SearchCasesByFields(l.ctx, builder, in.ProblemId, in.TestGroup)
	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, xcode.RecordNotFound
		}
		logx.Errorf("find testcases fail, err : %v, result : %+v", err, result)
		return nil, xcode.ServerErr
	}

	var testcases []*pb.Testcases
	for _, v := range result {
		testcases = append(testcases, &pb.Testcases{
			TestId:         v.TestId,
			ProblemId:      v.ProblemId,
			TestGroup:      v.TestGroup,
			InputFileName:  v.InputFilePath,
			OutputFileName: v.OutputFilePath,
			UpdateAt:       &timestamppb.Timestamp{Seconds: v.UpdateTime.Unix()},
		})
	}
	return &pb.SearchTestcasesResp{Testcases: testcases}, nil
}
