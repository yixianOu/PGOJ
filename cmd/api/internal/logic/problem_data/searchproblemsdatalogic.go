package problem_data

import (
	"context"
	"oj-micro/problems/cmd/rpc/pb"

	"oj-micro/problems/cmd/api/internal/svc"
	"oj-micro/problems/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchProblemsDataLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSearchProblemsDataLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchProblemsDataLogic {
	return &SearchProblemsDataLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SearchProblemsDataLogic) SearchProblemsData(req *types.SearchProblemsDataRequest) (resp *types.SearchProblemsDataResponse, err error) {
	result, err := l.svcCtx.ProblemServiceRpc.SearchProblemdata(l.ctx, &pb.SearchProblemdataReq{
		Page:       req.Page,
		Limit:      req.PageSize,
		ScoreFloor: req.ScoreFloor,
		ScoreCeil:  req.ScoreCeil,
		Auth:       req.Auth,
		Order:      req.Order,
	})
	if err != nil {
		return nil, err
	}

	problemsData := make([]types.ProblemData, 0)
	for _, v := range result.Problemdata {
		problemsData = append(problemsData, types.ProblemData{
			ProblemDataId:       v.ProblemdataId,
			ProblemId:           v.ProblemId,
			Submission:          v.Submission,
			Accepted:            v.Ac,
			MemoryLimitExceeded: v.Mle,
			TimeLimitExceeded:   v.Tle,
			WrongAnswer:         v.Wa,
			RuntimeError:        v.Rte,
			CompileError:        v.Ce,
			OutputLimitExceeded: v.Ole,
			UnknowError:         v.Ue,
			Score:               v.Score,
			Auth:                v.Auth,
			SegmentFault:        v.Sf,
			FloatError:          v.Fe,
		})
	}

	resp = &types.SearchProblemsDataResponse{
		ProblemData: problemsData,
	}
	return
}
