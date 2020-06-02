package rabbit

import "github.com/I-Reven/Hexagonal/src/domain/message/rabbit"

type Rabbit interface {
	Init(message rabbit.Message)
	Boot()
}
