package svc

import (
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/zrpc"
	"oj-micro/common/interceptors"
	"oj-micro/problems/cmd/rpc/problemservice"
	"oj-micro/users/cmd/rpc/userservice"

	//"github.com/zeromicro/go-zero/rest"
	"oj-micro/other/internal/config"
)

type ServiceContext struct {
	Config            config.Config
	MinioClients      []*minio.Client
	UserServiceRpc    userservice.UserService
	ProblemServiceRpc problemservice.ProblemService
}

func NewServiceContext(c config.Config) *ServiceContext {
	userRPC := zrpc.MustNewClient(c.UserRpcConf, zrpc.WithUnaryClientInterceptor(interceptors.ClientErrorInterceptor()))
	problemRPC := zrpc.MustNewClient(c.ProblemRpcConf, zrpc.WithUnaryClientInterceptor(interceptors.ClientErrorInterceptor()))
	return &ServiceContext{
		Config:            c,
		MinioClients:      NewOssClient(c),
		UserServiceRpc:    userservice.NewUserService(userRPC),
		ProblemServiceRpc: problemservice.NewProblemService(problemRPC),
	}
}

func NewOssClient(c config.Config) []*minio.Client {
	var minioClients []*minio.Client

	for _, minioConfig := range c.MinioConfig.Endpoints {
		minioClient, err := minio.New(minioConfig, &minio.Options{
			Creds:  credentials.NewStaticV4(c.MinioConfig.AccessKeyId, c.MinioConfig.AccessKeySecret, ""),
			Secure: false,
		})
		if err != nil || minioClient == nil {
			logx.Errorf("minio client init failed : %v", err)
			panic(err)
		}

		minioClients = append(minioClients, minioClient)
	}

	return minioClients
}
