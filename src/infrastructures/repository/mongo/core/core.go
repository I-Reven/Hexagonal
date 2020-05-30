package core

import (
	db "github.com/I-Reven/Hexagonal/src/infrastructures/repository/mongo"
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
	Core struct {
		DB db.Mongo
	}
)

func (c Core) Mongo() *Connection {
	once.Do(func() {
		connection = &Connection{db.Mongo{}.Connection(config)}
	})

	return connection
}
