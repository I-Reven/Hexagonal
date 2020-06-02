package core

import "github.com/I-Reven/Hexagonal/src/infrastructure/repository/mongo/core"

type Core interface {
	Connection() *core.Connection
}
