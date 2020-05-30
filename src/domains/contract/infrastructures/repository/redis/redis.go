package redis

import "time"

type Redis interface {
	Ping() (string, error)
	Get(key string) (string, error)
	Set(key string, value interface{}, expiration time.Duration) error
	Del(key string) error
}
