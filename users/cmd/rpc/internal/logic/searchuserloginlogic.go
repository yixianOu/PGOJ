package logic

import (
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"oj-micro/common/xcode"
	"oj-micro/users/cmd/rpc/internal/code"
	"oj-micro/users/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
	"oj-micro/users/cmd/rpc/internal/svc"
)

type SearchUserLoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchUserLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchUserLoginLogic {
	return &SearchUserLoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchUserLoginLogic) SearchUserLogin(in *pb.SearchUserLoginReq) (*pb.SearchUserLoginResp, error) {
	builder := l.svcCtx.UserLoginModel.SelectBuilder()
	loginRows, err := l.svcCtx.UserLoginModel.SearchUserByFields(l.ctx, builder, in.Page, in.Limit, in.RoleLevel, in.Username, in.Order)
	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, code.UserNotFoundError
		}
		logx.Errorf("from SearchUserLogin：UserLoginModel.FindListByPage失败:\n %v", err)
		return nil, xcode.ServerErr
	}

	var userLogin []*pb.UserLogin
	for _, loginRow := range loginRows {
		userLogin = append(userLogin, &pb.UserLogin{
			Id:            loginRow.Id,
			RoleLevel:     loginRow.RoleLevel,
			Username:      loginRow.Username,
			Email:         loginRow.Email,
			CoverImageUrl: loginRow.CoverImageUrl,
		})
	}

	return &pb.SearchUserLoginResp{UserLogin: userLogin}, nil
}
