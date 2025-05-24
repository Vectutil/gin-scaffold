package redis

import (
	"context"
	"fmt"
	"gin-scaffold/internal/config"
	"github.com/redis/go-redis/v9"
)

type redisClient struct {
	client *redis.Client
}

var rdb *redisClient

func Init() {
	// 初始化redis
	rdb = &redisClient{}
	rdb.client = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", config.Cfg.Redis.Host, config.Cfg.Redis.Port),
		Password: config.Cfg.Redis.Password, // no password set
		DB:       config.Cfg.Redis.DB,       // use default DB
	})
	_, err := rdb.client.Ping(context.Background()).Result()
	if err != nil {
		panic(err)
	}
}

func GetRedisClient() *redisClient {
	return rdb
}
