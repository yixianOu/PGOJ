package svc

import (
	"github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/zrpc"
	"oj-micro/common/interceptors"
	"oj-micro/users/cmd/api/internal/config"
	"oj-micro/users/cmd/api/internal/svc/email"
	"oj-micro/users/cmd/rpc/userservice"
)

type ServiceContext struct {
	Config         config.Config
	UserServiceRpc userservice.UserService
	//OssClient      *oss.Client
	Mailer *email.Email
	Redis  *redis.Client
}

//const (
//	defaultOssConnectTimeout   = 1
//	defaultOssReadWriteTimeout = 3
//)

func NewServiceContext(c config.Config) *ServiceContext {
	//if c.Oss.ConnectTimeout == 0 {
	//	c.Oss.ConnectTimeout = defaultOssConnectTimeout
	//}
	//if c.Oss.ReadWriteTimeout == 0 {
	//	c.Oss.ReadWriteTimeout = defaultOssReadWriteTimeout
	//}
	//oc, err := oss.New(c.Oss.Endpoint, c.Oss.AccessKeyId, c.Oss.AccessKeySecret,
	//	oss.Timeout(c.Oss.ConnectTimeout, c.Oss.ReadWriteTimeout))
	//if err != nil {
	//	panic(err)
	//}

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

	return &ServiceContext{
		Config:         c,
		UserServiceRpc: userservice.NewUserService(userRPC),
		//OssClient:      oc,
		Mailer: mailer,
		Redis:  rdb,
	}
}
