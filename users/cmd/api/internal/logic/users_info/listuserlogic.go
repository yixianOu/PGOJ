package users_info

import (
	"context"
	"oj-micro/users/cmd/rpc/pb"

	"oj-micro/users/cmd/api/internal/svc"
	"oj-micro/users/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListUserLogic {
	return &ListUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListUserLogic) ListUser(req *types.ListUserRequest) (resp *types.ListUserResponse, err error) {
	request := &pb.SearchUserLoginReq{
		Page:      int64(req.Page),
		Limit:     int64(req.PageSize),
		RoleLevel: req.RoleLevel,
		Username:  req.Username,
		Order:     req.Order,
	}

	searchUserLoginResp, err := l.svcCtx.UserServiceRpc.SearchUserLogin(l.ctx, request)
	if err != nil {
		return nil, err
	}

	var loginRows []types.LoginRow
	for _, loginRow := range searchUserLoginResp.UserLogin {
		loginRows = append(loginRows, types.LoginRow{
			ID:            loginRow.Id,
			RoleLevel:     loginRow.RoleLevel,
			Username:      loginRow.Username,
			CoverImageUrl: loginRow.CoverImageUrl,
		})
	}

	resp = &types.ListUserResponse{
		User_list:  loginRows,
		UsersCount: len(loginRows),
	}
	return
}
