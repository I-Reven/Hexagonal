package monngo

import "github.com/go-bongo/bongo"

type Mongo interface {
	Connection(config bongo.Config) *bongo.Connection
}
