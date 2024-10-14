package users_profile

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

type UpdateProfileLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateProfileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateProfileLogic {
	return &UpdateProfileLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateProfileLogic) UpdateProfile(req *types.UpdateProfileRequest) (resp *types.UpdateProfileResponse, err error) {
	//jwt是否过期
	expireTime, err := l.ctx.Value("access_expire").(json.Number).Int64()
	if err != nil || time.Now().Unix() > expireTime {
		return nil, xcode.UnauthorizedTokenTimeout
	}
	userID, err := l.ctx.Value("user_id").(json.Number).Int64()
	if err != nil || userID != req.UserID {
		return nil, xcode.UnauthorizedUserNotLogin
	}
	_, err = l.svcCtx.UserServiceRpc.UpdateUserProfile(l.ctx, &pb.UpdateUserProfileReq{
		UserId:      req.UserID,
		Phone:       req.Phone,
		Name:        req.Name,
		Description: req.Description,
		School:      req.School,
	})

	return
}
