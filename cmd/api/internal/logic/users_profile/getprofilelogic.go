package users_profile

import (
	"context"
	"oj-micro/users/cmd/rpc/pb"

	"oj-micro/users/cmd/api/internal/svc"
	"oj-micro/users/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetProfileLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetProfileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetProfileLogic {
	return &GetProfileLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetProfileLogic) GetProfile(req *types.GetProfileRequest) (resp *types.ProfileResponse, err error) {
	profileById, err := l.svcCtx.UserServiceRpc.GetUserProfileById(l.ctx, &pb.GetUserProfileByIdReq{
		UserId: req.UserID,
	})
	if err != nil {
		return nil, err
	}
	userLoginById, err := l.svcCtx.UserServiceRpc.GetUserLoginById(l.ctx, &pb.GetUserLoginByIdReq{
		Id: req.UserID,
	})
	if err != nil {
		return nil, err
	}

	resp = &types.ProfileResponse{
		Profile: types.Profile{
			UserID:      profileById.UserProfile.UserId,
			Phone:       profileById.UserProfile.Phone,
			Name:        profileById.UserProfile.Name,
			ACCount:     profileById.UserProfile.ACCount,
			SubmitCount: profileById.UserProfile.SubmitCount,
			Score:       profileById.UserProfile.Score,
			Description: profileById.UserProfile.Description,
			Rating:      profileById.UserProfile.Rating,
			ACProblem:   profileById.UserProfile.ACProblem,
			School:      profileById.UserProfile.School,
		},
		UserInfo: types.LoginRow{
			Username:      userLoginById.UserLogin.Username,
			ID:            userLoginById.UserLogin.Id,
			RoleLevel:     userLoginById.UserLogin.RoleLevel,
			CoverImageUrl: userLoginById.UserLogin.CoverImageUrl,
		},
	}
	return
}
