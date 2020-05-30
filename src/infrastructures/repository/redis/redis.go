package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"time"
)

type (
	Redis struct {
		Client *redis.Client
		Ctx    context.Context
	}
)

func (r *Redis) Ping() (string, error) {
	return r.Client.Ping(r.Ctx).Result()
}

func (r *Redis) Get(key string) (string, error) {
	return r.Client.Get(r.Ctx, key).Result()
}

func (r *Redis) Set(key string, value interface{}, expiration time.Duration) error {
	return r.Client.Set(r.Ctx, key, value, expiration).Err()
}

func (r *Redis) Del(key string) error {
	return r.Client.Del(r.Ctx, key).Err()
}
