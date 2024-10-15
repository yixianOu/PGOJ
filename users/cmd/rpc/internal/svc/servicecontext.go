package svc

import (
	"errors"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"log"
	"oj-micro/users/cmd/rpc/internal/config"
	"oj-micro/users/model"
)

type ServiceContext struct {
	Config      config.Config
	RedisClient *redis.Redis

	UserLoginModel   model.UserLoginModel
	UserProfileModel model.UserProfileModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	dbConn := sqlx.NewMysql(c.DB.DataSource)
	return &ServiceContext{
		Config:      c,
		RedisClient: redis.MustNewRedis(c.Redis.RedisConf),

		UserLoginModel:   model.NewUserLoginModel(dbConn, c.Cache),
		UserProfileModel: model.NewUserProfileModel(dbConn, c.Cache),
	}
}

func VerifyCode(rdb *redis.Redis, email, userInputCode string) bool {
	storedCode, err := rdb.Get(email)
	if errors.Is(err, redis.Nil) { // 如果键不存在，说明验证码已过期或从未设置
		return false
	} else if err != nil {
		log.Printf("获取验证码失败: %v", err)
		return false
	}
	return storedCode == userInputCode
}
