package rabbit

import (
	message "github.com/I-Reven/Hexagonal/src/domains/message/rabbit"
	"github.com/I-Reven/Hexagonal/src/infrastructures/logger"
	"github.com/juju/errors"
	"github.com/streadway/amqp"
)

func ConsumeMessage(message message.Message) (<-chan amqp.Delivery, error) {
	var err error
	var q amqp.Queue

	Init(message)

	q, err = ch.QueueDeclare(
		message.GetConsumerConfig().Name,           // name, leave empty to generate a unique name
		message.GetConsumerConfig().Durable,        // durable
		message.GetConsumerConfig().DeleteWhenUsed, // delete when used
		message.GetConsumerConfig().Exclusive,      // exclusive
		message.GetExchangeConfig().NoWait,         // noWait
		message.GetExchangeConfig().Arguments,      // arguments,   // arguments
	)

	if err != nil {
		err = errors.NewNotSupported(err, "error.rabbit-can-not-connect-to-server")
		logger.Error(err)
		return nil, err
	}

	err = ch.QueueBind(
		q.Name,                                 // name of the queue
		message.GetConsumerConfig().BindingKey, // bindingKey
		message.GetConsumerConfig().SourceExchange, // sourceExchange
		message.GetExchangeConfig().NoWait,         // noWait
		message.GetExchangeConfig().Arguments,      // arguments // arguments
	)

	if err != nil {
		err = errors.NewNotSupported(err, "error.rabbit-can-not-build-queue")
		logger.Error(err)
		return nil, err
	}

	return ch.Consume(
		q.Name,                                // queue
		message.GetConsumerConfig().Consumer,  // consumer
		message.GetConsumerConfig().AutoAck,   // auto-ack
		message.GetConsumerConfig().Exclusive, // exclusive
		message.GetConsumerConfig().NoLocal,   // no-local
		message.GetExchangeConfig().NoWait,    // noWait
		message.GetExchangeConfig().Arguments, // arguments
	)

}
