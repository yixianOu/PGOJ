package users_info

import (
	"context"
	"oj-micro/users/cmd/rpc/pb"

	"oj-micro/users/cmd/api/internal/svc"
	"oj-micro/users/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginWithCodeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginWithCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginWithCodeLogic {
	return &LoginWithCodeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginWithCodeLogic) LoginWithCode(req *types.LoginWithCodeRequest) (resp *types.LoginUserResponse, err error) {
	request := pb.LoginWithCodeReq{
		Email:     req.Email,
		EmailCode: req.Code,
	}

	loginUserResp, err := l.svcCtx.UserServiceRpc.LoginWithCode(l.ctx, &request)
	if err != nil {
		return nil, err
	}

	resp = &types.LoginUserResponse{
		User_login: types.LoginRow{
			ID:            loginUserResp.UserLogin.Id,
			RoleLevel:     loginUserResp.UserLogin.RoleLevel,
			Username:      loginUserResp.UserLogin.Username,
			CoverImageUrl: loginUserResp.UserLogin.CoverImageUrl,
		},
		Token: loginUserResp.Token,
	}

	return
}
