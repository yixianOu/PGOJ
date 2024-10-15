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

type UpdateUserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserInfoLogic {
	return &UpdateUserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateUserInfoLogic) UpdateUserInfo(req *types.UpdateUserRequest) (resp *types.UpdateUserResponse, err error) {
	expireTime, err := l.ctx.Value("access_expire").(json.Number).Int64()
	if err != nil || time.Now().Unix() > expireTime {
		return nil, xcode.UnauthorizedTokenTimeout
	}
	userID, err := l.ctx.Value("user_id").(json.Number).Int64()
	if err != nil || userID != req.ID {
		return nil, xcode.UnauthorizedUserNotLogin
	}

	_, err = l.svcCtx.UserServiceRpc.PartialUpdateUser(l.ctx, &pb.PartialUpdateUserReq{
		Id:            req.ID,
		Username:      req.Username,
		Email:         req.Email,
		CoverImageUrl: req.CoverImageUrl,
		Password:      req.Password,
	})

	return
}
