package logic

import (
	"context"
	"errors"
	"oj-micro/common/xcode"
	"oj-micro/users/cmd/rpc/internal/code"
	"oj-micro/users/cmd/rpc/pb"
	"oj-micro/users/model"

	"github.com/zeromicro/go-zero/core/logx"
	"oj-micro/users/cmd/rpc/internal/svc"
)

type GetUserLoginByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserLoginByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLoginByIdLogic {
	return &GetUserLoginByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserLoginByIdLogic) GetUserLoginById(in *pb.GetUserLoginByIdReq) (*pb.GetUserLoginByIdResp, error) {
	findOne, err := l.svcCtx.UserLoginModel.FindOne(l.ctx, in.Id)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return nil, code.UserNotFoundError
		} else {
			l.Logger.Errorf("查询用户失败: %v", err)
			return nil, xcode.ServerErr
		}
	}

	userLogin := &pb.UserLogin{
		Id:            findOne.Id,
		RoleLevel:     findOne.RoleLevel,
		Username:      findOne.Username,
		Email:         findOne.Email,
		CoverImageUrl: findOne.CoverImageUrl,
	}
	return &pb.GetUserLoginByIdResp{UserLogin: userLogin}, nil
}
