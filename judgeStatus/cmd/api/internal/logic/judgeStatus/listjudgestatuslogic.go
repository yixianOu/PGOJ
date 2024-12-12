package judgeStatus

import (
	"context"
	"encoding/json"
	"oj-micro/common/xcode"
	"oj-micro/judgeStatus/cmd/rpc/pb"
	"time"

	"oj-micro/judgeStatus/cmd/api/internal/svc"
	"oj-micro/judgeStatus/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListJudgeStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListJudgeStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListJudgeStatusLogic {
	return &ListJudgeStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListJudgeStatusLogic) ListJudgeStatus(req *types.ListJudgeStatusRequest) (resp *types.ListJudgeStatusResponse, err error) {
	expireTime, err := l.ctx.Value("access_expire").(json.Number).Int64()
	if err != nil || time.Now().Unix() > expireTime {
		return nil, xcode.UnauthorizedTokenTimeout
	}
	userID, err := l.ctx.Value("user_id").(json.Number).Int64()
	if err != nil || userID != req.UserId {
		return nil, xcode.AccessDenied
	}

	results, err := l.svcCtx.JudgeServiceRpc.SearchJudgestatus(l.ctx, &pb.SearchJudgestatusReq{
		Page:         req.Page,
		Limit:        req.Limit,
		UserId:       req.UserId,
		ProblemId:    req.ProblemId,
		Result:       req.Result,
		Language:     req.Language,
		SubmitTime:   req.SubmitTime,
		Contest:      req.Contest,
		ProblemTitle: req.ProblemTitle,
		Order:        req.Order,
	})
	if err != nil {
		return nil, err
	}

	judgeStatuses := make([]types.JudgeStatus, len(results.Judgestatus))
	for i, result := range results.Judgestatus {
		judgeStatuses[i] = types.JudgeStatus{
			JudgeId:        result.JudgeId,
			UserId:         result.UserId,
			ProblemId:      result.ProblemId,
			ProblemTitle:   result.ProblemTitle,
			Oj:             result.Oj,
			Result:         result.Result,
			TimeCost:       result.TimeCost,
			MemoryCost:     result.MemoryCost,
			Length:         result.Length,
			Language:       result.Language,
			SubmitTime:     result.SubmitTime,
			Code:           result.Code,
			Judger:         result.Judger,
			Contest:        result.Contest,
			ContestProblem: result.ContestProblem,
			TestCases:      result.TestCases,
			Message:        result.Message,
			Rating:         result.Rating,
			Ip:             result.Ip,
			InputData:      result.InputData,
			SampleOutput:   result.SampleOutput,
			UserOutput:     result.UserOutput,
		}
	}

	resp = &types.ListJudgeStatusResponse{
		JudgeStatus: judgeStatuses,
	}
	return
}
