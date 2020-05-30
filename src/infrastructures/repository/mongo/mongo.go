package mongo

import (
	"github.com/I-Reven/Hexagonal/src/infrastructures/logger"
	"github.com/go-bongo/bongo"
	"github.com/juju/errors"
)

type Mongo struct {
	Log logger.Log
}

func (m Mongo) Connection(config bongo.Config) *bongo.Connection {
	c, err := bongo.Connect(&config)

	if err != nil {
		err = errors.NewNotSupported(err, "error.mongo-can-not-connect-to-server")
		m.Log.Panic(err)
	}

	return c
}
