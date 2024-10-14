package logic

import (
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"oj-micro/common/xcode"
	"oj-micro/problems/cmd/rpc/internal/code"
	"oj-micro/problems/model"

	"oj-micro/problems/cmd/rpc/internal/svc"
	"oj-micro/problems/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddTagLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddTagLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddTagLogic {
	return &AddTagLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddTagLogic) AddTag(in *pb.AddTagReq) (*pb.AddTagResp, error) {
	oneByTagName, err := l.svcCtx.TagModel.FindOneByTagName(l.ctx, in.TagName)
	if !errors.Is(err, sqlx.ErrNotFound) {
		logx.Errorf("find tag by tag name fail, err : %v, oneByTagName : %+v", err, oneByTagName)
		return nil, code.TagExist
	}

	result, err := l.svcCtx.TagModel.Insert(l.ctx, &model.Tag{
		TagName: in.TagName,
	})
	if err != nil {
		logx.Errorf("insert tag fail, err : %v, result : %+v", err, result)
		return nil, xcode.ServerErr
	}
	return &pb.AddTagResp{}, nil
}
