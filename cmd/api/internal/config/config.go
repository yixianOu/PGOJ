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
	UserRpcConf zrpc.RpcClientConf
	//Oss         struct {
	//	Endpoint         string
	//	AccessKeyId      string
	//	AccessKeySecret  string
	//	BucketName       string
	//	AccessUrl        string
	//	ConnectTimeout   int64 `json:",optional"`
	//	ReadWriteTimeout int64 `json:",optional"`
	//}
	//UploadConfigs struct {
	//	ImageLimit struct {
	//		MaxSize   int64
	//		AllowExts []string
	//	}
	//}
	Email struct {
		Host     string
		Port     int
		Username string
		Password string
		IsSSL    bool
		From     string
	}
	Redis struct {
		Addr     string
		Password string
	}
}
