package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"oj-micro/common/interceptors"
	"oj-micro/problems/cmd/api/internal/config"
	"oj-micro/problems/cmd/rpc/problemservice"
	"oj-micro/users/cmd/rpc/userservice"
)

type ServiceContext struct {
	Config            config.Config
	ProblemServiceRpc problemservice.ProblemService
	UserServiceRpc    userservice.UserService
}

func NewServiceContext(c config.Config) *ServiceContext {
	problemRPC := zrpc.MustNewClient(c.ProblemRpcConf, zrpc.WithUnaryClientInterceptor(interceptors.ClientErrorInterceptor()))
	userRPC := zrpc.MustNewClient(c.UserRpcConf, zrpc.WithUnaryClientInterceptor(interceptors.ClientErrorInterceptor()))
	return &ServiceContext{
		Config:            c,
		ProblemServiceRpc: problemservice.NewProblemService(problemRPC),
		UserServiceRpc:    userservice.NewUserService(userRPC),
	}
}
