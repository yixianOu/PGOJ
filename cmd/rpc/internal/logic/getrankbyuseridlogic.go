package logic

import (
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"oj-micro/common/xcode"
	"oj-micro/users/model"

	"oj-micro/users/cmd/rpc/internal/svc"
	"oj-micro/users/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetRankByUserIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetRankByUserIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRankByUserIdLogic {
	return &GetRankByUserIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetRankByUserIdLogic) GetRankByUserId(in *pb.GetRankByUserIdReq) (*pb.GetRankByUserIdResp, error) {
	rank, err := l.svcCtx.UserProfileModel.SortUserByScoreAndReturnRank(l.ctx, in.UserId)
	if err != nil {
		logx.Errorf("from GetRankByUserId：SortUserByScoreAndReturnRank失败:\n %v", err)
		return nil, xcode.ServerErr
	}

	profile, err := l.svcCtx.UserProfileModel.FindOneByUserId(l.ctx, in.UserId)
	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, xcode.RecordNotFound
		}
		logx.Errorf("from GetRankByUserId：FindOneByUserId失败:\n %v", err)
		return nil, xcode.ServerErr
	}

	err = l.svcCtx.UserProfileModel.PartialUpdateProfile(l.ctx, &model.UserProfile{
		Id:     profile.Id,
		UserId: in.UserId,
		Rating: uint64(rank),
	})
	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, xcode.RecordNotFound
		}
		logx.Errorf("from GetRankByUserId：PartialUpdateProfile失败:\n %v", err)
		return nil, xcode.ServerErr
	}

	return &pb.GetRankByUserIdResp{UserProfile: &pb.UserProfile{
		Id:          profile.Id,
		UserId:      profile.UserId,
		Phone:       profile.Phone.String,
		Name:        profile.Name.String,
		Description: profile.Description.String,
		School:      profile.School.String,
		Score:       profile.Score,
		Rating:      rank,
		ACCount:     profile.ACCount,
		SubmitCount: profile.SubmitCount,
	}}, nil
}
