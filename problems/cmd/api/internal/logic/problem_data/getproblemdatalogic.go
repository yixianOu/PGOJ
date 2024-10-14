package problem_data

import (
	"context"
	"oj-micro/problems/cmd/rpc/pb"

	"oj-micro/problems/cmd/api/internal/svc"
	"oj-micro/problems/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetProblemDataLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetProblemDataLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetProblemDataLogic {
	return &GetProblemDataLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetProblemDataLogic) GetProblemData(req *types.GetProblemDataRequest) (resp *types.GetProblemDataResponse, err error) {
	result, err := l.svcCtx.ProblemServiceRpc.GetProblemdataById(l.ctx, &pb.GetProblemdataByIdReq{Id: req.ProblemDataId})
	if err != nil {
		return
	}

	resp = &types.GetProblemDataResponse{
		ProblemData: types.ProblemData{
			ProblemDataId:       result.Problemdata.ProblemdataId,
			ProblemId:           result.Problemdata.ProblemId,
			Submission:          result.Problemdata.Submission,
			Accepted:            result.Problemdata.Ac,
			MemoryLimitExceeded: result.Problemdata.Mle,
			TimeLimitExceeded:   result.Problemdata.Tle,
			WrongAnswer:         result.Problemdata.Wa,
			RuntimeError:        result.Problemdata.Rte,
			CompileError:        result.Problemdata.Ce,
			OutputLimitExceeded: result.Problemdata.Ole,
			UnknowError:         result.Problemdata.Ue,
			Score:               result.Problemdata.Score,
			Auth:                result.Problemdata.Auth,
			SegmentFault:        result.Problemdata.Sf,
			FloatError:          result.Problemdata.Fe,
		},
	}
	return
}
