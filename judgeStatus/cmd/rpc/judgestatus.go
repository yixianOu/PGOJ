package main

import (
	"flag"
	"fmt"
	"oj-micro/common/interceptors"

	"oj-micro/judgeStatus/cmd/rpc/internal/config"
	"oj-micro/judgeStatus/cmd/rpc/internal/server"
	"oj-micro/judgeStatus/cmd/rpc/internal/svc"
	"oj-micro/judgeStatus/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "judgeStatus/cmd/rpc/etc/judgestatus.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		pb.RegisterJudgeServiceServer(grpcServer, server.NewJudgeServiceServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	s.AddUnaryInterceptors(interceptors.ServerErrorInterceptor())
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
