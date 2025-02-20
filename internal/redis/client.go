package redis

import (
	"codegen-service/internal/config"
	"context"
	"github.com/go-redis/redis/v8"
	"time"
)

var ctx = context.Background()

type RedisClient struct {
	client *redis.Client
}

func NewRedisClient(cfg *config.Config) *RedisClient {
	client := redis.NewClient(&redis.Options{
		Addr: cfg.RedisAddress,
		DB:   cfg.RedisDb,
	})
	return &RedisClient{client: client}
}

func (r *RedisClient) SetData(key, value string, duration *time.Duration) {
	if duration == nil {
		defaultDuration := 20 * time.Minute
		duration = &defaultDuration
	}
	r.client.Set(ctx, key, value, *duration)
}

func (r *RedisClient) GetData(key string) (string, error) {
	val, err := r.client.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}
	return val, nil
}
