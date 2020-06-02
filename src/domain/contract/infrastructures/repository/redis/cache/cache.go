package cache

import "github.com/I-Reven/Hexagonal/src/infrastructure/repository/redis"

type Cache interface {
	Init() redis.Redis
}
