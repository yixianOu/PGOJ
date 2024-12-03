package logic

import (
	"context"
	"errors"
	"golang.org/x/crypto/bcrypt"
	app "oj-micro/common/jwt"
	"oj-micro/common/xcode"
	"oj-micro/users/cmd/rpc/internal/code"
	"oj-micro/users/model"

	"oj-micro/users/cmd/rpc/internal/svc"
	"oj-micro/users/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginUserLogic {
	return &LoginUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginUserLogic) LoginUser(in *pb.LoginUserReq) (*pb.LoginUserResp, error) {
	oneByEmail, err := l.svcCtx.UserLoginModel.FindOneByEmail(l.ctx, in.Email)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return nil, code.UserEmailNotExist
		} else {
			l.Logger.Errorf("from LoginUser：UserLoginModel.FindOneByEmail失败:\n %v", err)
			return nil, xcode.ServerErr
		}
	}

	err = bcrypt.CompareHashAndPassword([]byte(oneByEmail.Password), []byte(in.Password))
	if err != nil {
		return nil, code.PasswordError
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
	if err != nil {
		l.Logger.Errorf("from LoginUser：app.GenerateToken失败:\n %v", err)
		return nil, xcode.ServerErr
	}

	return &pb.LoginUserResp{Token: token, UserLogin: UserLogin}, nil
}
