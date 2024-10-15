package email

import (
	"context"
	"github.com/redis/go-redis/v9"
	"oj-micro/users/cmd/api/internal/config"
	"time"
)

var ctx = context.Background()

func NewRedisClient(RedisSetting *config.Config) (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     RedisSetting.Redis.Addr,
		Password: RedisSetting.Redis.Password,
	})
	_, err := rdb.Ping(ctx).Result()
	return rdb, err
}

func SaveCode(rdb *redis.Client, email, code string) error {
	err := rdb.SetEx(ctx, email, code, 5*time.Minute).Err()
	return err
}
