package logic

import (
	"context"
	"oj-micro/common/xcode"
	"oj-micro/users/cmd/rpc/pb"
	"oj-micro/users/model"

	"github.com/zeromicro/go-zero/core/logx"
	"oj-micro/users/cmd/rpc/internal/svc"
)

type UpdateUserProfileLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateUserProfileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserProfileLogic {
	return &UpdateUserProfileLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateUserProfileLogic) UpdateUserProfile(in *pb.UpdateUserProfileReq) (*pb.UpdateUserProfileResp, error) {
	data := &model.UserProfile{
		Id:          in.Id,
		UserId:      in.UserId,
		ACCount:     in.ACCount,
		SubmitCount: in.SubmitCount,
		Score:       in.Score,
		Rating:      uint64(in.Rating),
	}

	if in.Phone == "" {
		data.Phone.Valid = false
	} else {
		data.Phone.String = in.Phone
		data.Phone.Valid = true
	}
	if in.Name == "" {
		data.Name.Valid = false
	} else {
		data.Name.String = in.Name
		data.Name.Valid = true
	}
	if in.Description == "" {
		data.Description.Valid = false
	} else {
		data.Description.String = in.Description
		data.Description.Valid = true
	}
	if in.School == "" {
		data.School.Valid = false
	} else {
		data.School.String = in.School
		data.School.Valid = true
	}

	//if in.ACCount == 0 {
	//	data.ACCount = userProfile.ACCount
	//} else {
	//	data.ACCount = in.ACCount
	//}
	//if in.SubmitCount == 0 {
	//	data.SubmitCount = userProfile.SubmitCount
	//} else {
	//	data.SubmitCount = in.SubmitCount
	//}
	//if in.Score == 0 {
	//	data.Score = userProfile.Score
	//} else {
	//	data.Score = in.Score
	//}
	//if in.Rating == 0 {
	//	data.Rating = userProfile.Rating
	//} else {
	//	data.Rating = uint64(in.Rating)
	//}

	err := l.svcCtx.UserProfileModel.PartialUpdateProfile(l.ctx, data)
	if err != nil {
		l.Logger.Errorf("from UpdateUserProfile：Update失败:\n %v", err)
		return nil, xcode.ServerErr
	}
	return &pb.UpdateUserProfileResp{}, nil
}
