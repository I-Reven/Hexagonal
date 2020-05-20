package core

import (
	"github.com/I-Reven/Hexagonal/src/infrastructures/repository/mongo"
	"github.com/go-bongo/bongo"
	"os"
	"sync"
)

var (
	config = bongo.Config{
		ConnectionString: os.Getenv("MONGO_URL"),
		Database:         "core",
	}
	once       sync.Once
	connection *Connection
)

type (
	Connection struct {
		*bongo.Connection
	}
)

func Mongo() *Connection {
	once.Do(func() {
		connection = &Connection{mongo.Init(config)}
	})

	return connection
}
