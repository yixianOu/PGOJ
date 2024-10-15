package users_info

import (
	"context"
	"encoding/json"
	"oj-micro/common/xcode"
	"oj-micro/users/cmd/rpc/pb"
	"time"

	"oj-micro/users/cmd/api/internal/svc"
	"oj-micro/users/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteUserLogic {
	return &DeleteUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteUserLogic) DeleteUser(req *types.DeleteUserRequest) (resp *types.DeleteUserResponse, err error) {
	expireTime, err := l.ctx.Value("access_expire").(json.Number).Int64()
	if err != nil || time.Now().Unix() > expireTime {
		return nil, xcode.UnauthorizedTokenTimeout
	}
	userID, err := l.ctx.Value("user_id").(json.Number).Int64()
	if err != nil || userID != req.ID {
		return nil, xcode.UnauthorizedUserNotLogin
	}

	_, err = l.svcCtx.UserServiceRpc.DelUserLogin(l.ctx, &pb.DelUserLoginReq{Id: req.ID})
	return
}
