package problems

import (
	"context"
	"oj-micro/problems/cmd/rpc/pb"

	"oj-micro/problems/cmd/api/internal/svc"
	"oj-micro/problems/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchProblemsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSearchProblemsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchProblemsLogic {
	return &SearchProblemsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SearchProblemsLogic) SearchProblems(req *types.SearchProblemsRequest) (resp *types.SearchProblemsResponse, err error) {
	result, err := l.svcCtx.ProblemServiceRpc.SearchProblem(l.ctx, &pb.SearchProblemReq{
		Page:        req.Page,
		Limit:       req.PageSize,
		Author:      req.Author,
		Title:       req.Title,
		ProblemCode: req.ProblemCode,
		Oj:          req.Oj,
		Des:         req.Description,
		Source:      req.Source,
		Auth:        req.Auth,
		Level:       req.Level,
		Order:       req.Order,
	})
	if err != nil {
		return nil, err
	}

	problems := make([]types.Problem, len(result.Problem))
	for i, problem := range result.Problem {
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

	resp = &types.SearchProblemsResponse{
		Problems: problems,
	}
	return
}
