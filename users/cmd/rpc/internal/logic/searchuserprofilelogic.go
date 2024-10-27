package logic

import (
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"oj-micro/common/xcode"

	"oj-micro/users/cmd/rpc/internal/svc"
	"oj-micro/users/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchUserProfileLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchUserProfileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchUserProfileLogic {
	return &SearchUserProfileLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchUserProfileLogic) SearchUserProfile(in *pb.SearchUserProfileReq) (*pb.SearchUserProfileResp, error) {
	builder := l.svcCtx.UserProfileModel.SelectBuilder()
	profiles, err := l.svcCtx.UserProfileModel.SearchUserProfileByFields(l.ctx,
		builder, in.Page, in.Limit, in.Description, in.School, in.OrderByScore)
	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, xcode.RecordNotFound
		}
		logx.Errorf("from SearchUserProfile：SearchUserProfileByFields失败:\n %v", err)
		return nil, xcode.ServerErr
	}

	userProfiles := make([]*pb.UserProfile, 0)
	for _, profile := range profiles {
		userProfiles = append(userProfiles, &pb.UserProfile{
			Id:          profile.Id,
			UserId:      profile.UserId,
			Phone:       profile.Phone.String,
			Name:        profile.Name.String,
			Description: profile.Description.String,
			School:      profile.School.String,
			Score:       profile.Score,
			ACCount:     profile.ACCount,
			SubmitCount: profile.SubmitCount,
			Rating:      int64(profile.Rating),
		})
	}

	return &pb.SearchUserProfileResp{UserProfile: userProfiles}, nil
}
