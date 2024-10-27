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
	UserRpcConf        zrpc.RpcClientConf
	JudgeStatusRpcConf zrpc.RpcClientConf
	ProblemRpcConf     zrpc.RpcClientConf

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
