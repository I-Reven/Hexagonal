package mongo

import (
	"github.com/I-Reven/Hexagonal/infrastructures/logger"
	"github.com/go-bongo/bongo"
)

func Init(config bongo.Config) *bongo.Connection {
		c, err := bongo.Connect(&config)

		logger.LOG().FatalE(err)

	return c
}
