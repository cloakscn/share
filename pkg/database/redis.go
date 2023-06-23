package database

import (
	"cloaks.cn/share/pkg/configuration"
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

func redisOpen(cfg configuration.RedisConfiguration) *redis.Client {
	dsn := fmt.Sprintf("%s:%s",
		cfg.Host,
		cfg.Port,
	)

	client := redis.NewClient(&redis.Options{
		Addr:     dsn,
		DB:       0,
		Username: cfg.UserName,
		Password: cfg.Password,
	})

	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		panic("failed to connect redis")
	}

	return client
}
