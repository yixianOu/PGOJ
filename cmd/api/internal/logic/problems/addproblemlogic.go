package problems

import (
	"context"
	"encoding/json"
	"oj-micro/common/dataType"
	"oj-micro/common/xcode"
	"oj-micro/problems/cmd/rpc/pb"
	userpb "oj-micro/users/cmd/rpc/pb"
	"time"

	"oj-micro/problems/cmd/api/internal/svc"
	"oj-micro/problems/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddProblemLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddProblemLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddProblemLogic {
	return &AddProblemLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddProblemLogic) AddProblem(req *types.AddProblemRequest) (resp *types.AddProblemResponse, err error) {
	expireTime, err := l.ctx.Value("access_expire").(json.Number).Int64()
	if err != nil || time.Now().Unix() > expireTime {
		return nil, xcode.UnauthorizedTokenTimeout
	}
	userID, err := l.ctx.Value("user_id").(json.Number).Int64()
	if err != nil {
		return nil, xcode.UnauthorizedUserNotLogin
	}
	roleLevel, err := l.ctx.Value("role_level").(json.Number).Int64()
	userLevel := dataType.RoleLevel(roleLevel)
	if err != nil || userLevel != dataType.AdminUser {
		return nil, xcode.UnauthorizedUserNotSuperUser
	}

	result, err := l.svcCtx.UserServiceRpc.GetUserLoginById(l.ctx, &userpb.GetUserLoginByIdReq{
		Id: userID,
	})
	if err != nil {
		return nil, err
	}

	_, err = l.svcCtx.ProblemServiceRpc.AddProblem(l.ctx, &pb.AddProblemReq{
		Title:       req.Title,
		Author:      result.UserLogin.Username,
		Oj:          req.Oj,
		Des:         req.Description,
		Input:       req.Input,
		Output:      req.Output,
		Hint:        req.Hint,
		Source:      req.Source,
		Time:        req.LimitTime,
		Memory:      req.LimitMemory,
		Auth:        req.Auth,
		Level:       req.Level,
		ProblemCode: req.ProblemCode,

		Score:  req.Score,
		TagIds: req.TagIds,
	})
	return
}
