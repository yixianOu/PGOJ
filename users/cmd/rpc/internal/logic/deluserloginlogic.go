package logic

import (
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/logx"
	"oj-micro/common/xcode"
	"oj-micro/users/cmd/rpc/internal/code"
	"oj-micro/users/cmd/rpc/internal/svc"
	"oj-micro/users/cmd/rpc/pb"
	"oj-micro/users/model"
)

type DelUserLoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelUserLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelUserLoginLogic {
	return &DelUserLoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelUserLoginLogic) DelUserLogin(in *pb.DelUserLoginReq) (*pb.DelUserLoginResp, error) {
	userProfile, err := l.svcCtx.UserProfileModel.FindOneByUserId(l.ctx, in.Id)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return nil, code.UserNotFoundError
		} else {
			l.Logger.Errorf("from DelUserLogin：UserProfileModel.FindOneByUserId失败:\n %v", err)
			return nil, xcode.ServerErr
		}
	}
	err = l.svcCtx.UserProfileModel.Delete(l.ctx, userProfile.Id)
	if err != nil {
		l.Logger.Errorf("from DelUserLogin：UserProfileModel.Delete失败:\n %v", err)
		return nil, xcode.ServerErr
	}

	err = l.svcCtx.UserLoginModel.Delete(l.ctx, in.Id)
	if err != nil {
		l.Logger.Errorf("from DelUserLogin：UserLoginModel.Delete失败:\n %v", err)
		return nil, xcode.ServerErr
	}

	return &pb.DelUserLoginResp{}, nil
}
