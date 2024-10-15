package tags

import (
	"context"
	"oj-micro/problems/cmd/rpc/pb"

	"oj-micro/problems/cmd/api/internal/svc"
	"oj-micro/problems/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListTagsByProblemIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListTagsByProblemIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListTagsByProblemIdLogic {
	return &ListTagsByProblemIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListTagsByProblemIdLogic) ListTagsByProblemId(req *types.ListTagsByProblemIdRequest) (resp *types.ListTagsByProblemIdResponse, err error) {
	result, err := l.svcCtx.ProblemServiceRpc.ListTagsByProblemId(l.ctx, &pb.ListTagsByProblemIdReq{
		ProblemId: req.ProblemId,
	})
	if err != nil {
		return nil, err
	}

	tags := make([]types.Tag, len(result.Tags))
	for i, tag := range result.Tags {
		tags[i] = types.Tag{
			TagId:   tag.TagId,
			TagName: tag.TagName,
		}
	}

	resp = &types.ListTagsByProblemIdResponse{
		Tags: tags,
	}
	return
}
