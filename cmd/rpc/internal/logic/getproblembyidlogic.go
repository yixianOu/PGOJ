package logic

import (
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"google.golang.org/protobuf/types/known/timestamppb"
	"oj-micro/common/xcode"
	"oj-micro/problems/cmd/rpc/internal/svc"
	"oj-micro/problems/cmd/rpc/pb"
)

type GetProblemByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetProblemByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetProblemByIdLogic {
	return &GetProblemByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetProblemByIdLogic) GetProblemById(in *pb.GetProblemByIdReq) (*pb.GetProblemByIdResp, error) {
	result, err := l.svcCtx.ProblemModel.FindOne(l.ctx, in.Id)
	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, xcode.RecordNotFound
		}
		logx.Errorf("find problem by id fail, err : %v, result : %+v", err, result)
		return nil, xcode.ServerErr
	}

	return &pb.GetProblemByIdResp{Problem: &pb.Problem{
		ProblemId: result.ProblemId,
		Author:    result.Author,
		Addtime:   &timestamppb.Timestamp{Seconds: result.Addtime.Unix()},
		Oj:        result.Oj,
		Title:     result.Title,
		Des:       result.Des,
		Input:     result.Input,
		Output:    result.Output,
		Sinput:    result.Sinput,
		Soutput:   result.Soutput,
		Hint:      result.Hint,
		Source:    result.Source,
		Time:      result.Time,
		Memory:    result.Memory,
		Auth:      result.Auth,
		Level:     result.Level,
		//TestCount:   result.TestCount,
		ProblemCode: result.ProblemCode,
	}}, nil
}
