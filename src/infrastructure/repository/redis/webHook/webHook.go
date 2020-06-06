package webHook

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	repository "github.com/I-Reven/Hexagonal/src/infrastructure/repository/redis"
	redisV8 "github.com/go-redis/redis/v8"
	"github.com/juju/errors"
	"os"
	"sync"
	"time"
)

var (
	once       sync.Once
	redis      repository.Redis
	expiration = 1440 * time.Minute
	config     = &redisV8.Options{
		Addr:     os.Getenv("REDIS_URL"),
		Password: "",
		DB:       3,
	}
)

type WebHook struct{}

func (r *WebHook) redis() *repository.Redis {
	once.Do(func() {
		redis = repository.Redis{Client: redisV8.NewClient(config), Ctx: context.Background()}
	})

	return &redis
}

func (r *WebHook) MackKey(data []byte) string {
	hash := md5.Sum(data)
	return hex.EncodeToString(hash[:])
}

func (r *WebHook) GetKey(jobName string, hash string) string {
	return jobName + ":" + hash
}

func (r *WebHook) Create(jobName string, data []byte) (string, error) {
	key := r.GetKey(jobName, r.MackKey(data))
	return r.MackKey(data), r.Save(key, data)
}

func (r *WebHook) Get(jobName string, hash string) ([]byte, error) {
	data, err := r.redis().Get(r.GetKey(jobName, hash))

	if err != nil {
		err = errors.NewNotFound(err, "error.can-not-found-wen-hook-data")
		return nil, err
	}

	return []byte(data), nil
}

func (r *WebHook) Delete(jobName string, hash string) error {
	err := r.redis().Del(r.GetKey(jobName, hash))

	if err != nil {
		err = errors.NewNotFound(err, "error.can-not-found-wen-hook-data")
	}

	return err
}

func (r *WebHook) Save(key string, data []byte) error {
	return r.redis().Set(key, data, expiration)
}
