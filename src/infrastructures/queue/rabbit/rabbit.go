package rabbit

import (
	"flag"
	"github.com/I-Reven/Hexagonal/src/domains/message/rabbit"
	"github.com/I-Reven/Hexagonal/src/infrastructures/logger"
	"github.com/juju/errors"
	"github.com/streadway/amqp"
	"os"
	"sync"
)

var (
	amqpURI *string
	conn    *amqp.Connection
	ch      *amqp.Channel
	once    sync.Once
)

func Boot() {
	once.Do(func() {
		var err error
		amqpURI = flag.String("amqp", os.Getenv("RABBIT_URL"), "AMQP URI")
		flag.Parse()
		conn, err = amqp.Dial(*amqpURI)

		if err != nil {
			err = errors.NewNotSupported(err, "error.rabbit-can-not-connect-to-server")
			logger.Panic(err)
		}

	})
}

func Init(message rabbit.Message) {
	Boot()
	var err error
	ch, err = conn.Channel()

	if err != nil {
		err = errors.NewNotSupported(err, "error.rabbit-can-not-connect-to-channel")
		logger.Panic(err)
	}

	err = ch.ExchangeDeclare(
		message.GetExchangeConfig().Name,        // name
		message.GetExchangeConfig().Type,        // type
		message.GetExchangeConfig().Durable,     // durable
		message.GetExchangeConfig().AutoDeleted, // auto-deleted
		message.GetExchangeConfig().Internal,    // internal
		message.GetExchangeConfig().NoWait,      // noWait
		message.GetExchangeConfig().Arguments,   // arguments
	)

	if err != nil {
		err = errors.NewNotSupported(err, "error.rabbit-can-not-exchange-declare-channel")
		logger.Panic(err)
	}
}
