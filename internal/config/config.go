package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	Auth struct {
		AccessSecret string
		AccessExpire int64
	}
	UploadConfigs struct {
		ImageLimit struct {
			MaxSize   int64
			AllowExts []string
		}
		SampleIOLimit struct {
			MaxSize   int64
			AllowExts []string
		}
	}

	MinioConfig struct {
		Endpoints       []string
		AccessKeyId     string
		AccessKeySecret string
		BucketName      string
		DomainName      string
	}
	//Email struct {
	//	Host     string
	//	Port     int
	//	Username string
	//	Password string
	//	IsSSL    bool
	//	From     string
	//}
	UserRpcConf    zrpc.RpcClientConf
	ProblemRpcConf zrpc.RpcClientConf
}
