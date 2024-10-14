package problems

import (
	"context"
	"oj-micro/problems/cmd/rpc/pb"

	"oj-micro/problems/cmd/api/internal/svc"
	"oj-micro/problems/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListProblemsByTagIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListProblemsByTagIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListProblemsByTagIdLogic {
	return &ListProblemsByTagIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListProblemsByTagIdLogic) ListProblemsByTagId(req *types.ListProblemsByTagIdRequest) (resp *types.ListProblemsByTagIdResponse, err error) {
	result, err := l.svcCtx.ProblemServiceRpc.ListProblemsByTagId(l.ctx, &pb.ListProblemsByTagIdReq{
		TagId: req.TagId,
		Page:  req.Page,
		Limit: req.PageSize,
	})
	if err != nil {
		return nil, err
	}

	problems := make([]types.Problem, len(result.Problems))
	for i, problem := range result.Problems {
		problems[i] = types.Problem{
			ProblemId:    problem.ProblemId,
			Author:       problem.Author,
			Addtime:      problem.Addtime.AsTime().Unix(),
			Oj:           problem.Oj,
			Title:        problem.Title,
			Description:  problem.Des,
			Input:        problem.Input,
			Output:       problem.Output,
			SampleInput:  problem.Sinput,
			SampleOutput: problem.Soutput,
			Hint:         problem.Hint,
			Source:       problem.Source,
			LimitTime:    problem.Time,
			LimitMemory:  problem.Memory,
			Auth:         problem.Auth,
			Level:        problem.Level,
			//Test_count:   problem.TestCount,
			ProblemCode: problem.ProblemCode,
		}
	}
	resp = &types.ListProblemsByTagIdResponse{
		Problems: problems,
	}
	return
}
