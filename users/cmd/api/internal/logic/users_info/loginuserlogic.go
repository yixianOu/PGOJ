package users_info

import (
	"context"
	"oj-micro/users/cmd/rpc/pb"

	"oj-micro/users/cmd/api/internal/svc"
	"oj-micro/users/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginUserLogic {
	return &LoginUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginUserLogic) LoginUser(req *types.LoginUserRequest) (resp *types.LoginUserResponse, err error) {
	request := pb.LoginUserReq{
		Email:    req.Email,
		Password: req.Password,
	}

	loginUserResp, err := l.svcCtx.UserServiceRpc.LoginUser(l.ctx, &request)
	if err != nil {
		return nil, err
	}

	resp = &types.LoginUserResponse{
		Token: loginUserResp.Token,
		User_login: types.LoginRow{
			ID:            loginUserResp.UserLogin.Id,
			RoleLevel:     loginUserResp.UserLogin.RoleLevel,
			Username:      loginUserResp.UserLogin.Username,
			CoverImageUrl: loginUserResp.UserLogin.CoverImageUrl,
		},
	}

	return
}
