package tags

import (
	"context"
	"oj-micro/problems/cmd/rpc/pb"

	"oj-micro/problems/cmd/api/internal/svc"
	"oj-micro/problems/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchTagsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSearchTagsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchTagsLogic {
	return &SearchTagsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SearchTagsLogic) SearchTags(req *types.SearchTagRequest) (resp *types.SearchTagResponse, err error) {
	result, err := l.svcCtx.ProblemServiceRpc.SearchTag(l.ctx, &pb.SearchTagReq{
		Page:    req.Page,
		Limit:   req.Limit,
		TagName: req.TagName,
		Order:   req.Order,
	})
	if err != nil {
		return nil, err
	}

	tags := make([]types.Tag, len(result.Tag))
	for i, tag := range result.Tag {
		tags[i] = types.Tag{
			TagId:   tag.TagId,
			TagName: tag.TagName,
		}
	}

	resp = &types.SearchTagResponse{
		Tag: tags,
	}
	return
}
