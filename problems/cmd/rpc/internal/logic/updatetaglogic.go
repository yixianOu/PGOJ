package logic

import (
	"context"
	"errors"
	"oj-micro/common/xcode"
	"oj-micro/problems/cmd/rpc/internal/code"
	"oj-micro/problems/model"

	"oj-micro/problems/cmd/rpc/internal/svc"
	"oj-micro/problems/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateTagLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateTagLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateTagLogic {
	return &UpdateTagLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateTagLogic) UpdateTag(in *pb.UpdateTagReq) (*pb.UpdateTagResp, error) {
	newData := &model.Tag{
		TagId:   in.TagId,
		TagName: in.TagName,
	}

	oneByTagName, err := l.svcCtx.TagModel.FindOneByTagName(l.ctx, newData.TagName)
	if !errors.Is(err, model.ErrNotFound) && oneByTagName.TagId != newData.TagId {
		return nil, code.TagNameExist
	}

	err = l.svcCtx.TagModel.Update(l.ctx, newData)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return nil, xcode.RecordNotFound
		}
		l.Logger.Errorf("update tag fail, err : %v, newData : %+v", err, newData)
		return nil, xcode.ServerErr
	}
	return &pb.UpdateTagResp{}, nil
}
