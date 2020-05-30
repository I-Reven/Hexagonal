package session

import (
	"github.com/I-Reven/Hexagonal/src/infrastructures/logger"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/juju/errors"
	"os"
)

type Session struct {
	log logger.Log
}

func (s *Session) Store() sessions.Store {
	store, err := redis.NewStore(10, "tcp", os.Getenv("REDIS_URL"), "", []byte("secret"))

	if err != nil {
		err = errors.NewNotSupported(err, "error.session-can-not-connect-to-redis")
		s.log.Error(err)
	}

	return store
}
