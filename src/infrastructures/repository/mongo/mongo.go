package mongo

import (
	"github.com/I-Reven/Hexagonal/src/infrastructures/logger"
	"github.com/go-bongo/bongo"
	"github.com/juju/errors"
)

func Init(config bongo.Config) *bongo.Connection {
	c, err := bongo.Connect(&config)

	if err != nil {
		err = errors.NewNotSupported(err, "error.mongo-can-not-connect-to-server")
		logger.Panic(err)
	}

	return c
}
