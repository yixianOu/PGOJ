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

type GetUserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserInfoLogic) GetUserInfo(req *types.GetUserInfoRequest) (resp *types.GetUserInfoResponse, err error) {
	expireTime, err := l.ctx.Value("access_expire").(json.Number).Int64()
	if err != nil || time.Now().Unix() > expireTime {
		return nil, xcode.UnauthorizedTokenTimeout
	}
	userID, err := l.ctx.Value("user_id").(json.Number).Int64()
	if err != nil || userID != req.ID {
		return nil, xcode.UnauthorizedUserNotLogin
	}

	byIdResp, err := l.svcCtx.UserServiceRpc.GetUserLoginById(l.ctx, &pb.GetUserLoginByIdReq{Id: req.ID})
	if err != nil {
		return nil, err
	}

	resp = &types.GetUserInfoResponse{
		User_info: types.UserInfo{
			ID:            byIdResp.UserLogin.Id,
			Username:      byIdResp.UserLogin.Username,
			RoleLevel:     byIdResp.UserLogin.RoleLevel,
			CoverImageUrl: byIdResp.UserLogin.CoverImageUrl,
			Email:         byIdResp.UserLogin.Email,
			Password:      "",
		},
	}

	return
}
