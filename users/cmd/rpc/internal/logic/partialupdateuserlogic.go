package logic

import (
	"context"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"oj-micro/common/xcode"
	"oj-micro/users/cmd/rpc/internal/code"
	"oj-micro/users/model"

	"oj-micro/users/cmd/rpc/internal/svc"
	"oj-micro/users/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type PartialUpdateUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPartialUpdateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PartialUpdateUserLogic {
	return &PartialUpdateUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PartialUpdateUserLogic) PartialUpdateUser(in *pb.PartialUpdateUserReq) (*pb.PartialUpdateUserResp, error) {
	var hashedPassword []byte
	err := error(nil)

	if in.Password != "" {
		hashedPassword, err = bcrypt.GenerateFromPassword([]byte(in.Password), bcrypt.DefaultCost)
		if err != nil {
			l.Logger.Errorf("bcrypt.GenerateFromPassword error: %v", err)
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
		if errors.Is(err, model.ErrNotFound) {
			return nil, code.UserNotFoundError
		}
		l.Logger.Errorf("PartialUpdateUser error: %v,Value: %v", err, data)
		return nil, xcode.ServerErr
	}

	return &pb.PartialUpdateUserResp{}, nil
}
