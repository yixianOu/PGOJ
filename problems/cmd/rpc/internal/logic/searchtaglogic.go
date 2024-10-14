package logic

import (
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"oj-micro/common/xcode"
	"oj-micro/problems/cmd/rpc/internal/svc"
	"oj-micro/problems/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchTagLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchTagLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchTagLogic {
	return &SearchTagLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchTagLogic) SearchTag(in *pb.SearchTagReq) (*pb.SearchTagResp, error) {
	builder := l.svcCtx.TagModel.SelectBuilder()
	result, err := l.svcCtx.TagModel.SearchTagsByFields(l.ctx, builder, in.Page, in.Limit, in.TagName, in.Order)
	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, xcode.RecordNotFound
		}
		logx.Errorf("search tags by fields fail, err : %v, result : %+v", err, result)
		return nil, xcode.ServerErr
	}

	tags := make([]*pb.Tag, 0)
	for _, v := range result {
		tags = append(tags, &pb.Tag{
			TagId:   v.TagId,
			TagName: v.TagName,
		})
	}
	return &pb.SearchTagResp{Tag: tags}, nil
}
