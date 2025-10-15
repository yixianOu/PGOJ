package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"oj-micro/common/interceptors"
	"oj-micro/judgeStatus/cmd/api/internal/config"
	"oj-micro/judgeStatus/cmd/rpc/judgeservice"
	"oj-micro/problems/cmd/rpc/problemservice"
)

type ServiceContext struct {
	Config            config.Config
	JudgeServiceRpc   judgeservice.JudgeService
	ProblemServiceRpc problemservice.ProblemService
}

func NewServiceContext(c config.Config) *ServiceContext {
	judgeRPC := zrpc.MustNewClient(c.JudgeStatusRpcConf, zrpc.WithUnaryClientInterceptor(interceptors.ClientErrorInterceptor()))
	problemRPC := zrpc.MustNewClient(c.ProblemRpcConf, zrpc.WithUnaryClientInterceptor(interceptors.ClientErrorInterceptor()))
	return &ServiceContext{
		Config:            c,
		JudgeServiceRpc:   judgeservice.NewJudgeService(judgeRPC),
		ProblemServiceRpc: problemservice.NewProblemService(problemRPC),
	}
}
