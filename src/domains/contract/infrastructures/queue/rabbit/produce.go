package rabbit

import message "github.com/I-Reven/Hexagonal/src/domains/message/rabbit"

type Produce interface {
	ProduceMessage(message message.Message) error
}
