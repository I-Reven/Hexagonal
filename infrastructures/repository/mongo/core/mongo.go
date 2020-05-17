package core

import (
	"github.com/go-bongo/bongo"
	"log"
	"sync"
)

var(
	config = &bongo.Config{
		ConnectionString: "localhost",
		Database:         "core",
	}
	once sync.Once
	connection *Connection
)

type(
	Connection struct {
		*bongo.Connection
	}
)

func Mongo() *Connection {
	once.Do(func() {
		c, err := bongo.Connect(config)
		if err != nil {
			log.Fatal(err)
		}
		connection = &Connection{c}
	})

	return connection
}
