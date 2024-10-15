package problems

import (
	"context"
	"oj-micro/problems/cmd/rpc/pb"

	"oj-micro/problems/cmd/api/internal/svc"
	"oj-micro/problems/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetProblemByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetProblemByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetProblemByIdLogic {
	return &GetProblemByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetProblemByIdLogic) GetProblemById(req *types.GetProblemByIdRequest) (resp *types.GetProblemByIdResponse, err error) {
	result, err := l.svcCtx.ProblemServiceRpc.GetProblemById(l.ctx, &pb.GetProblemByIdReq{
		Id: req.ProblemId,
	})
	if err != nil {
		return nil, err
	}

	resp = &types.GetProblemByIdResponse{
		Problem: types.Problem{
			ProblemId:    result.Problem.ProblemId,
			Author:       result.Problem.Author,
			Addtime:      result.Problem.Addtime.AsTime().Unix(),
			Oj:           result.Problem.Oj,
			Title:        result.Problem.Title,
			Description:  result.Problem.Des,
			Input:        result.Problem.Input,
			Output:       result.Problem.Output,
			SampleInput:  result.Problem.Sinput,
			SampleOutput: result.Problem.Soutput,
			Hint:         result.Problem.Hint,
			Source:       result.Problem.Source,
			LimitTime:    result.Problem.Time,
			LimitMemory:  result.Problem.Memory,
			Auth:         result.Problem.Auth,
			Level:        result.Problem.Level,
			//Test_count:   result.Problem.TestCounts,
			ProblemCode: result.Problem.ProblemCode,
		},
	}
	return
}
