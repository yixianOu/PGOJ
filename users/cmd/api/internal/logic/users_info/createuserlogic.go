package users_info

import (
	"context"
	"oj-micro/users/cmd/rpc/pb"

	"oj-micro/users/cmd/api/internal/svc"
	"oj-micro/users/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateUserLogic {
	return &CreateUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateUserLogic) CreateUser(req *types.CreateUserRequest) (resp *types.CreateUserResponse, err error) {
	request := pb.AddUserLoginReq{
		Password:      req.Password,
		Username:      req.Username,
		Email:         req.Email,
		CoverImageUrl: req.CoverImageUrl,
		EmailCode:     req.EmailCode,
	}
	_, err = l.svcCtx.UserServiceRpc.AddUserLogin(l.ctx, &request)

	return
}
