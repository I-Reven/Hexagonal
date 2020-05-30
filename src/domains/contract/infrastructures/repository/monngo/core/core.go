package core

import "github.com/I-Reven/Hexagonal/src/infrastructures/repository/mongo/core"

type Core interface {
	Connection() *core.Connection
}
