package users_info

import (
	"context"
	"encoding/json"
	"oj-micro/common/dataType"
	"oj-micro/common/xcode"
	"oj-micro/users/cmd/rpc/pb"
	"time"

	"oj-micro/users/cmd/api/internal/svc"
	"oj-micro/users/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AuthenticateUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAuthenticateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AuthenticateUserLogic {
	return &AuthenticateUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AuthenticateUserLogic) AuthenticateUser(req *types.AuthorizeUserRequest) (resp *types.AuthorizeUserResponse, err error) {
	expireTime, err := l.ctx.Value("access_expire").(json.Number).Int64()
	if err != nil || time.Now().Unix() > expireTime {
		return nil, xcode.UnauthorizedTokenTimeout
	}
	roleLevel, err := l.ctx.Value("role_level").(json.Number).Int64()
	_roleLevel := dataType.RoleLevel(roleLevel)
	if err != nil || _roleLevel != dataType.AdminUser {
		return nil, xcode.UnauthorizedUserNotSuperUser
	}

	_, err = l.svcCtx.UserServiceRpc.PartialUpdateUser(l.ctx, &pb.PartialUpdateUserReq{
		Id:        req.UserID,
		RoleLevel: req.RoleLevel,
	})

	return
}
