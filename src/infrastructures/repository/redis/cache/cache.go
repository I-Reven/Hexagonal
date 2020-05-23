package cache

import (
	"context"
	"github.com/I-Reven/Hexagonal/src/infrastructures/repository/redis"
	redisV8 "github.com/go-redis/redis/v8"
	"os"
	"sync"
)

var (
	once   sync.Once
	client *redisV8.Client
	ctx    context.Context
	config = &redisV8.Options{
		Addr:     os.Getenv("REDIS_URL"),
		Password: "",
		DB:       0,
	}
)

func Cache() redis.Redis {
	once.Do(func() {
		client = redisV8.NewClient(config)
		ctx = context.Background()
	})

	return redis.Redis{Client: client, Ctx: ctx}
}
