package logic

import (
	"context"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"oj-micro/common/xcode"
	"oj-micro/users/cmd/rpc/internal/code"
	"oj-micro/users/cmd/rpc/pb"
	"oj-micro/users/model"

	"github.com/zeromicro/go-zero/core/logx"
	"oj-micro/users/cmd/rpc/internal/svc"
)

type AddUserLoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddUserLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddUserLoginLogic {
	return &AddUserLoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// AddUserLogin 新用户权限默认为1
func (l *AddUserLoginLogic) AddUserLogin(in *pb.AddUserLoginReq) (*pb.AddUserLoginResp, error) {
	_, err := l.svcCtx.UserLoginModel.FindOneByEmail(l.ctx, in.Email)
	if !errors.Is(err, model.ErrNotFound) {
		return nil, code.EmailDuplicate
	}
	//验证验证码
	isValidated := svc.VerifyCode(l.svcCtx.RedisClient, in.Email, in.EmailCode)
	if !isValidated {
		return nil, code.EmailCodeError
	}
	//验证用户名是否重复
	_, err = l.svcCtx.UserLoginModel.FindOneByUsername(l.ctx, in.Username)
	if !errors.Is(err, model.ErrNotFound) {
		return nil, code.UserNameDuplicate
	}

	hashedPassword, erro := bcrypt.GenerateFromPassword([]byte(in.Password), bcrypt.DefaultCost)
	if erro != nil {
		l.Logger.Errorf("bcrypt.GenerateFromPassword error: %v", erro)
		return nil, xcode.ServerErr
	}
	userLogin := &model.UserLogin{
		Password:      string(hashedPassword),
		Username:      in.Username,
		Email:         in.Email,
		CoverImageUrl: in.CoverImageUrl,
		RoleLevel:     1,
	}
	result, err := l.svcCtx.UserLoginModel.Insert(l.ctx, userLogin)
	if err != nil {
		l.Logger.Errorf("UserLoginModel.Insert error: %v", err)
		return nil, xcode.ServerErr
	}
	userID, err := result.LastInsertId()

	_, err = l.svcCtx.UserProfileModel.Insert(l.ctx, &model.UserProfile{
		UserId: userID,
	})
	if err != nil {
		l.Logger.Errorf("UserProfileModel.Insert error: %v", err)
		return nil, xcode.ServerErr
	}

	return &pb.AddUserLoginResp{}, nil
}
