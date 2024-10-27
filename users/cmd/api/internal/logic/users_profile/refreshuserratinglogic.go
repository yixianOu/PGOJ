package users_profile

import (
	"context"
	"encoding/json"
	"oj-micro/common/dataType"
	"oj-micro/common/xcode"
	judgePb "oj-micro/judgeStatus/cmd/rpc/pb"
	problemPb "oj-micro/problems/cmd/rpc/pb"
	"oj-micro/users/cmd/api/internal/svc"
	"oj-micro/users/cmd/api/internal/types"
	"oj-micro/users/cmd/rpc/pb"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type RefreshUserRatingLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// NewRefreshUserRatingLogic 刷新用户判题数据
func NewRefreshUserRatingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RefreshUserRatingLogic {
	return &RefreshUserRatingLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RefreshUserRatingLogic) RefreshUserRating(req *types.RefreshUserRatingRequest) (resp *types.RefreshUserRatingResponse, err error) {
	expireTime, err := l.ctx.Value("access_expire").(json.Number).Int64()
	if err != nil || time.Now().Unix() > expireTime {
		return nil, xcode.UnauthorizedTokenTimeout
	}
	userID, err := l.ctx.Value("user_id").(json.Number).Int64()
	if err != nil || userID != req.UserID {
		return nil, xcode.UnauthorizedUserNotLogin
	}

	judges, err := l.svcCtx.JudgeServiceRpc.SearchJudgestatus(l.ctx, &judgePb.SearchJudgestatusReq{
		UserId: userID,
	})
	if err != nil {
		return nil, err
	}

	var acCount, score int64
	var firstAcProblemMap = make(map[int64]struct{})

	for _, judge := range judges.Judgestatus {
		if judge.Result == string(dataType.Accept) {
			acCount++
			//如果是第一次AC，则加分
			if _, ok := firstAcProblemMap[judge.ProblemId]; !ok {
				problemData, err2 := l.svcCtx.ProblemServiceRpc.GetProblemdataByProblemId(l.ctx, &problemPb.GetProblemdataByProblemIdReq{
					ProblemId: judge.ProblemId,
				})
				if err2 != nil {
					return nil, err2
				}
				firstAcProblemMap[judge.ProblemId] = struct{}{}
				score += problemData.Problemdata.Score
			}
		}
	}

	_, err = l.svcCtx.UserServiceRpc.UpdateUserProfile(l.ctx, &pb.UpdateUserProfileReq{
		UserId:      userID,
		SubmitCount: int64(len(judges.Judgestatus)),
		ACCount:     acCount,
		Score:       score,
	})
	if err != nil {
		return nil, err
	}

	profile, err := l.svcCtx.UserServiceRpc.GetRankByUserId(l.ctx, &pb.GetRankByUserIdReq{
		UserId: userID,
	})
	if err != nil {
		return nil, err
	}

	resp = &types.RefreshUserRatingResponse{
		Profile: types.Profile{
			ID:          profile.UserProfile.Id,
			UserID:      profile.UserProfile.UserId,
			Phone:       profile.UserProfile.Phone,
			Name:        profile.UserProfile.Name,
			Description: profile.UserProfile.Description,
			School:      profile.UserProfile.School,
			Score:       profile.UserProfile.Score,
			Rating:      profile.UserProfile.Rating,
			ACCount:     profile.UserProfile.ACCount,
			SubmitCount: profile.UserProfile.SubmitCount,
		},
	}

	return
}
