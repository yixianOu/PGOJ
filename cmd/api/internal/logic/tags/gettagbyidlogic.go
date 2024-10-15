package tags

import (
	"context"
	"oj-micro/problems/cmd/rpc/pb"

	"oj-micro/problems/cmd/api/internal/svc"
	"oj-micro/problems/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTagByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetTagByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTagByIdLogic {
	return &GetTagByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetTagByIdLogic) GetTagById(req *types.GetTagByIdRequest) (resp *types.GetTagByIdResponse, err error) {
	result, err := l.svcCtx.ProblemServiceRpc.GetTagById(l.ctx, &pb.GetTagByIdReq{
		Id: req.TagId,
	})
	if err != nil {
		return nil, err
	}

	resp = &types.GetTagByIdResponse{
		Tag: types.Tag{
			TagId:   result.Tag.TagId,
			TagName: result.Tag.TagName,
		},
	}

	return
}
