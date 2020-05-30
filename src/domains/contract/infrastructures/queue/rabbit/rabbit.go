package rabbit

import "github.com/I-Reven/Hexagonal/src/domains/message/rabbit"

type Rabbit interface {
	Init(message rabbit.Message)
	Boot()
}
