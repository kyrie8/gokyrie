package conf

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
)

var rdClient *redis.Client
var Duration = 30 * 24 * 60 * 60 * time.Second
var ctx = context.Background()

type RedisClient struct {
}

func InitRedis() (*RedisClient, error) {
	rdClient = redis.NewClient(&redis.Options{
		Addr:     viper.GetString("redis.url"),
		Password: viper.GetString("redis.password"),
		DB:       0,
	})
	_, err := rdClient.Ping(context.Background()).Result()
	if err != nil {
		return nil, err
	}
	return &RedisClient{}, nil
}

func (rc *RedisClient) Set(key string, value any) error {
	return rdClient.Set(ctx, key, value, Duration).Err()
}

func (rc *RedisClient) Get(key string) (any, error) {
	return rdClient.Get(ctx, key).Result()
}

func (rc *RedisClient) Delete(key ...string) error {
	return rdClient.Del(ctx, key...).Err()
}
