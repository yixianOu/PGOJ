package logic

import (
	"context"
	"errors"
	"oj-micro/common/xcode"
	"oj-micro/users/cmd/rpc/internal/code"
	"oj-micro/users/cmd/rpc/pb"
	"oj-micro/users/model"

	"github.com/zeromicro/go-zero/core/logx"
	"oj-micro/users/cmd/rpc/internal/svc"
)

type GetUserProfileByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserProfileByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserProfileByIdLogic {
	return &GetUserProfileByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserProfileByIdLogic) GetUserProfileById(in *pb.GetUserProfileByIdReq) (*pb.GetUserProfileByIdResp, error) {
	userProfile, err := l.svcCtx.UserProfileModel.FindOneByUserId(l.ctx, in.UserId)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return nil, code.UserNotFoundError
		} else {
			logx.Errorf("from GetUserProfileById：FindOneByUserId失败:\n %v", err)
			return nil, xcode.ServerErr
		}
	}
	userProfileResp := &pb.UserProfile{
		Id:          userProfile.Id,
		UserId:      userProfile.UserId,
		Phone:       userProfile.Phone.String,
		Name:        userProfile.Name.String,
		ACCount:     userProfile.ACCount,
		SubmitCount: userProfile.SubmitCount,
		Score:       userProfile.Score,
		Description: userProfile.Description.String,
		Rating:      int64(userProfile.Rating),
		ACProblem:   userProfile.ACProblem.String,
		School:      userProfile.School.String,
	}
	return &pb.GetUserProfileByIdResp{UserProfile: userProfileResp}, nil
}
