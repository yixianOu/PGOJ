package logic

import (
	"context"
	"database/sql"
	"errors"
	"oj-micro/common/xcode"
	"oj-micro/users/cmd/rpc/internal/code"
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
	userProfile, err := l.svcCtx.UserProfileModel.FindOneByUserId(l.ctx, in.UserId)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return nil, code.UserNotFoundError
		} else {
			logx.Errorf("from UpdateUserProfile：FindOneByUserId失败:\n %v", err)
			return nil, xcode.ServerErr
		}
	}

	var phone, name, description, school sql.NullString
	if in.Phone == "" {
		phone.Valid = false
	} else {
		phone.String = in.Phone
		phone.Valid = true
	}
	if in.Name == "" {
		name.Valid = false
	} else {
		name.String = in.Name
		name.Valid = true
	}
	if in.Description == "" {
		description.Valid = false
	} else {
		description.String = in.Description
		description.Valid = true
	}
	if in.School == "" {
		school.Valid = false
	} else {
		school.String = in.School
		school.Valid = true
	}

	data := &model.UserProfile{
		Id:          userProfile.Id,
		UserId:      in.UserId,
		Phone:       phone,
		Name:        name,
		ACCount:     in.ACCount,
		SubmitCount: in.SubmitCount,
		Score:       in.Score,
		Rating:      uint64(in.Rating),
		Description: description,
		School:      school,
	}
	err = l.svcCtx.UserProfileModel.Update(l.ctx, data)
	if err != nil {
		logx.Errorf("from UpdateUserProfile：Update失败:\n %v", err)
		return nil, xcode.ServerErr
	}
	return &pb.UpdateUserProfileResp{}, nil
}
