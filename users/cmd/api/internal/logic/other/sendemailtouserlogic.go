package other

import (
	"context"
	"fmt"
	"math/rand"
	"oj-micro/common/xcode"
	"oj-micro/users/cmd/api/internal/svc/email"

	"oj-micro/users/cmd/api/internal/svc"
	"oj-micro/users/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendEmailToUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSendEmailToUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendEmailToUserLogic {
	return &SendEmailToUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SendEmailToUserLogic) SendEmailToUser(req *types.SendEmailToUserRequest) (resp *types.SendEmailToUserResponse, err error) {
	//if req.Email == "" {
	//	return nil, code.EmailFormatError
	//}
	//随机生成验证码，发给用户邮箱，并添加到Context中
	emailCode := fmt.Sprintf("%06d", rand.Intn(1000000))
	subject := "OJ，用户邮箱验证码"
	body := "您的验证码是：" + emailCode + "，有效期为5分钟。"

	err = l.svcCtx.Mailer.SendMail(req.Email, subject, body)
	if err != nil {
		l.Logger.Errorf("send email to %s failed: %s", req.Email, err)
		return nil, xcode.ServerErr
	}

	err = email.SaveCode(l.svcCtx.Redis, req.Email, emailCode)
	if err != nil {
		l.Logger.Errorf("save email_code to redis failed: %s", err)
		return nil, xcode.ServerErr
	}
	return
}
