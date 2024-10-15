package logic

import (
	"context"
	"errors"
	app "oj-micro/common/jwt"
	"oj-micro/common/xcode"
	"oj-micro/users/cmd/rpc/internal/code"
	"oj-micro/users/model"

	"oj-micro/users/cmd/rpc/internal/svc"
	"oj-micro/users/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginWithCodeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginWithCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginWithCodeLogic {
	return &LoginWithCodeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginWithCodeLogic) LoginWithCode(in *pb.LoginWithCodeReq) (*pb.LoginUserResp, error) {
	//验证验证码
	isValidated := svc.VerifyCode(l.svcCtx.RedisClient, in.Email, in.EmailCode)
	if !isValidated {
		return nil, code.EmailCodeError
	}
	//根据邮箱查找用户
	oneByEmail, err := l.svcCtx.UserLoginModel.FindOneByEmail(l.ctx, in.Email)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return nil, code.UserNotFoundError
		} else {
			logx.Errorf("from LoginWithCode：UserLoginModel.FindOneByEmail失败:\n %v", err)
			return nil, xcode.ServerErr
		}
	}
	UserLogin := &pb.UserLogin{
		Id:            oneByEmail.Id,
		RoleLevel:     oneByEmail.RoleLevel,
		Username:      oneByEmail.Username,
		Email:         oneByEmail.Email,
		CoverImageUrl: oneByEmail.CoverImageUrl,
	}

	//生成token
	token, err := app.GenerateToken(app.TokenOptions{
		UserID:       oneByEmail.Id,
		RoleLevel:    oneByEmail.RoleLevel,
		AccessExpire: l.svcCtx.Config.JWT.AccessExpire,
		AccessSecret: l.svcCtx.Config.JWT.AccessSecret,
	})
	return &pb.LoginUserResp{UserLogin: UserLogin, Token: token}, nil
}
