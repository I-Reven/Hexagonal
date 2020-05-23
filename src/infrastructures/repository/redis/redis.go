package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"time"
)

type (
	//Redis Struct
	Redis struct {
		Client *redis.Client
		Ctx    context.Context
	}
)

// Ping Redis
func (r Redis) Ping() (string, error) {
	return r.Client.Ping(r.Ctx).Result()
}

// Get Value
func (r Redis) Get(key string) (string, error) {
	return r.Client.Get(r.Ctx, key).Result()
}

//Set Value
func (r Redis) Set(key string, value interface{}, expiration time.Duration) (string, error) {
	return r.Client.Set(r.Ctx, key, value, expiration).Result()
}
