package svc

import (
	"github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/zrpc"
	"oj-micro/common/interceptors"
	"oj-micro/judgeStatus/cmd/rpc/judgeservice"
	"oj-micro/problems/cmd/rpc/problemservice"
	"oj-micro/users/cmd/api/internal/config"
	"oj-micro/users/cmd/api/internal/svc/email"
	"oj-micro/users/cmd/rpc/userservice"
)

type ServiceContext struct {
	Config            config.Config
	UserServiceRpc    userservice.UserService
	JudgeServiceRpc   judgeservice.JudgeService
	ProblemServiceRpc problemservice.ProblemService
	Mailer            *email.Email
	Redis             *redis.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	mailer := email.NewEmail(&email.SMTPInfo{
		Host:     c.Email.Host,
		Port:     c.Email.Port,
		UserName: c.Email.Username,
		Password: c.Email.Password,
		IsSSL:    c.Email.IsSSL,
		From:     c.Email.From,
	})

	rdb, err := email.NewRedisClient(&c)
	if err != nil {
		panic(err)
	}

	userRPC := zrpc.MustNewClient(c.UserRpcConf, zrpc.WithUnaryClientInterceptor(interceptors.ClientErrorInterceptor()))
	judgeRPC := zrpc.MustNewClient(c.JudgeStatusRpcConf, zrpc.WithUnaryClientInterceptor(interceptors.ClientErrorInterceptor()))
	problemRPC := zrpc.MustNewClient(c.ProblemRpcConf, zrpc.WithUnaryClientInterceptor(interceptors.ClientErrorInterceptor()))

	return &ServiceContext{
		Config:            c,
		UserServiceRpc:    userservice.NewUserService(userRPC),
		JudgeServiceRpc:   judgeservice.NewJudgeService(judgeRPC),
		ProblemServiceRpc: problemservice.NewProblemService(problemRPC),
		Mailer:            mailer,
		Redis:             rdb,
	}
}
