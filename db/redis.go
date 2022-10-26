package db

import (
	"context"
	"fmt"
	"github.com/charlie-bit/yanxue/config"
	"github.com/go-redis/redis/v8"
)

var RedisClient *redis.Client

func InitRedis(cfg *config.Config) error {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     cfg.Redis.Addr,
		Username: cfg.Redis.Username,
		Password: cfg.Redis.Password,
		DB:       cfg.Redis.DB,
	})

	_, err := RedisClient.Ping(context.Background()).Result()
	if err != nil {
		return fmt.Errorf("redis connect failed  {err: %v}", err.Error())
	}

	return nil
}
