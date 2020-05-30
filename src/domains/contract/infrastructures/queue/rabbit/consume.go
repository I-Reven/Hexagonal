package rabbit

import (
	message "github.com/I-Reven/Hexagonal/src/domains/message/rabbit"
	"github.com/streadway/amqp"
)

type Consume interface {
	Message(message message.Message) (<-chan amqp.Delivery, error)
}
