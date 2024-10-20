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

type SearchProblemLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchProblemLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchProblemLogic {
	return &SearchProblemLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchProblemLogic) SearchProblem(in *pb.SearchProblemReq) (*pb.SearchProblemResp, error) {
	builder := l.svcCtx.ProblemModel.SelectBuilder()
	result, err := l.svcCtx.ProblemModel.SearchProblemByFields(l.ctx, builder, in.Page, in.Limit, in.Author, in.Oj, in.Title, in.ProblemCode, in.Des, in.Source, in.Auth, in.Level, in.Order)
	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, xcode.RecordNotFound
		}
		logx.Errorf("search problem by fields fail, err : %v, result : %+v", err, result)
		return nil, xcode.ServerErr
	}

	problem := make([]*pb.Problem, 0)
	for _, v := range result {
		problem = append(problem, &pb.Problem{
			ProblemId: v.ProblemId,
			Author:    v.Author,
			Addtime:   &timestamppb.Timestamp{Seconds: v.CreateTime.Unix()},
			Oj:        v.Oj,
			Title:     v.Title,
			Des:       v.Des,
			Input:     v.Input,
			Output:    v.Output,
			Sinput:    v.Sinput,
			Soutput:   v.Soutput,
			Hint:      v.Hint,
			Source:    v.Source,
			Time:      v.Time,
			Memory:    v.Memory,
			Auth:      v.Auth,
			Level:     v.Level,
			//TestCount:   v.TestCount,
			ProblemCode: v.ProblemCode,
		})
	}
	return &pb.SearchProblemResp{Problem: problem}, nil
}
