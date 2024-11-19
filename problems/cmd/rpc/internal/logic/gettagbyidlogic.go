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

type GetTagByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetTagByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTagByIdLogic {
	return &GetTagByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetTagByIdLogic) GetTagById(in *pb.GetTagByIdReq) (*pb.GetTagByIdResp, error) {
	result, err := l.svcCtx.TagModel.FindOne(l.ctx, in.Id)
	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, xcode.RecordNotFound
		}
		logx.Errorf("find tag by id fail, err : %v, result : %+v", err, result)
		return nil, xcode.ServerErr
	}
	return &pb.GetTagByIdResp{Tag: &pb.Tag{
		TagId:   result.TagId,
		TagName: result.TagName,
	}}, nil
}