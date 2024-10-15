package logic

import (
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"golang.org/x/crypto/bcrypt"
	"oj-micro/common/xcode"
	"oj-micro/users/cmd/rpc/internal/code"
	"oj-micro/users/cmd/rpc/pb"
	"oj-micro/users/model"

	"github.com/zeromicro/go-zero/core/logx"
	"oj-micro/users/cmd/rpc/internal/svc"
)

type UpdateUserLoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateUserLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserLoginLogic {
	return &UpdateUserLoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateUserLoginLogic) UpdateUserLogin(in *pb.UpdateUserLoginReq) (*pb.UpdateUserLoginResp, error) {
	var hashedPassword []byte
	err := error(nil)

	if in.Password != "" {
		hashedPassword, err = bcrypt.GenerateFromPassword([]byte(in.Password), bcrypt.DefaultCost)
		if err != nil {
			logx.Errorf("bcrypt.GenerateFromPassword error: %v", err)
			return nil, xcode.ServerErr
		}
	}
	data := &model.UserLogin{
		Id:            in.Id,
		RoleLevel:     in.RoleLevel,
		Username:      in.Username,
		Email:         in.Email,
		CoverImageUrl: in.CoverImageUrl,
		Password:      string(hashedPassword),
	}
	err = l.svcCtx.UserLoginModel.PartialUpdate(l.ctx, data)
	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, code.UserNotFoundError
		}
		logx.Errorf("UpdateUserLogin error: %v", err)
		return nil, xcode.ServerErr
	}
	return &pb.UpdateUserLoginResp{}, nil
}
