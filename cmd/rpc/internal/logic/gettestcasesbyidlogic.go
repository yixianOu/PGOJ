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

type GetTestcasesByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetTestcasesByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTestcasesByIdLogic {
	return &GetTestcasesByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetTestcasesByIdLogic) GetTestcasesById(in *pb.GetTestcasesByIdReq) (*pb.GetTestcasesByIdResp, error) {
	result, err := l.svcCtx.TestCasesModel.FindOne(l.ctx, in.Id)
	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, xcode.RecordNotFound
		}
		logx.Errorf("find testcases fail, err : %v, result : %+v", err, result)
		return nil, xcode.ServerErr
	}

	return &pb.GetTestcasesByIdResp{Testcases: &pb.Testcases{
		TestId:         result.TestId,
		ProblemId:      result.ProblemId,
		TestGroup:      result.TestGroup,
		InputFileName:  result.InputFilePath,
		OutputFileName: result.OutputFilePath,
		UpdateAt:       &timestamppb.Timestamp{Seconds: result.UpdateTime.Unix()},
	}}, nil
}
