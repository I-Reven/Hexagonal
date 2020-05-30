package cache

import "github.com/I-Reven/Hexagonal/src/infrastructures/repository/redis"

type Cache interface {
	Init() redis.Redis
}
