package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

type RedisClient struct {
	client *redis.Client
}

func NewRedisClient(addr string, db int) *RedisClient {
	client := redis.NewClient(&redis.Options{
		Addr: addr,
		DB:   db,
	})
	return &RedisClient{client: client}
}

func (r *RedisClient) Close() error {
	return r.client.Close()
}

func (r *RedisClient) SetData(key, value string) error {
	err := r.client.Set(ctx, key, value, 0).Err()
	if err != nil {
		return err
	}
	return nil
}

func (r *RedisClient) GetData(key string) (string, error) {
	val, err := r.client.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}
	return val, nil
}
