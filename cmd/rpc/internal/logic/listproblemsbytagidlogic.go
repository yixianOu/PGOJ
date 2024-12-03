package logic

import (
	"context"
	"errors"
	"google.golang.org/protobuf/types/known/timestamppb"
	"oj-micro/common/xcode"
	"oj-micro/problems/cmd/rpc/internal/svc"
	"oj-micro/problems/cmd/rpc/pb"
	"oj-micro/problems/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListProblemsByTagIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListProblemsByTagIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListProblemsByTagIdLogic {
	return &ListProblemsByTagIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ListProblemsByTagIdLogic) ListProblemsByTagId(in *pb.ListProblemsByTagIdReq) (*pb.ListProblemsByTagIdResp, error) {
	builder := l.svcCtx.ProblemTagModel.SelectBuilder()
	result, err := l.svcCtx.ProblemTagModel.FindProblemsByTagId(l.ctx, builder, in.TagId, in.Page, in.Limit)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return nil, xcode.RecordNotFound
		}
		l.Logger.Errorf("find problems by tag id fail, err : %v, result : %+v", err, result)
		return nil, xcode.ServerErr
	}

	problems := make([]*pb.Problem, 0)
	for _, v := range result {
		problem, err := l.svcCtx.ProblemModel.FindOne(l.ctx, v.ProblemId)
		if err != nil {
			if errors.Is(err, model.ErrNotFound) {
				return nil, xcode.RecordNotFound
			}
			l.Logger.Errorf("find problem by id fail, err : %v, result : %+v", err, problem)
			return nil, xcode.ServerErr
		}
		problems = append(problems, &pb.Problem{
			ProblemId: problem.ProblemId,
			Author:    problem.Author,
			Addtime:   &timestamppb.Timestamp{Seconds: problem.CreateTime.Unix()},
			Oj:        problem.Oj,
			Title:     problem.Title,
			Des:       problem.Des,
			Input:     problem.Input,
			Output:    problem.Output,
			Sinput:    problem.Sinput,
			Soutput:   problem.Soutput,
			Hint:      problem.Hint,
			Source:    problem.Source,
			Time:      problem.Time,
			Memory:    problem.Memory,
			Auth:      problem.Auth,
			Level:     problem.Level,
			//TestCount:   problem.TestCount,
			ProblemCode: problem.ProblemCode,
		})
	}
	return &pb.ListProblemsByTagIdResp{Problems: problems}, nil
}
