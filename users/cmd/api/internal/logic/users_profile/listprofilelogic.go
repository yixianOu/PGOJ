package users_profile

import (
	"context"
	"oj-micro/users/cmd/rpc/pb"

	"oj-micro/users/cmd/api/internal/svc"
	"oj-micro/users/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListProfileLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// NewListProfileLogic 列出用户判题信息
func NewListProfileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListProfileLogic {
	return &ListProfileLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListProfileLogic) ListProfile(req *types.ListProfileRequest) (resp *types.ListProfileResponse, err error) {
	profiles, err := l.svcCtx.UserServiceRpc.SearchUserProfile(l.ctx, &pb.SearchUserProfileReq{
		Page:         req.Page,
		Limit:        req.Limit,
		Description:  req.Descrition,
		School:       req.School,
		OrderByScore: req.OrderByScore,
	})
	if err != nil {
		return nil, err
	}

	var userProfiles []types.Profile
	for _, profile := range profiles.UserProfile {
		userProfiles = append(userProfiles, types.Profile{
			ID:          profile.Id,
			UserID:      profile.UserId,
			Phone:       profile.Phone,
			Name:        profile.Name,
			Description: profile.Description,
			School:      profile.School,
			Score:       profile.Score,
			ACCount:     profile.ACCount,
			SubmitCount: profile.SubmitCount,
			Rating:      profile.Rating,
		})
	}

	resp = &types.ListProfileResponse{
		Profiles: userProfiles,
	}

	return
}
