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

type Rabbit struct {
	Log logger.Log
}

func (r Rabbit) Init(message rabbit.Message) {
	r.Boot()
	var err error
	ch, err = conn.Channel()

	if err != nil {
		err = errors.NewNotSupported(err, "error.rabbit-can-not-connect-to-channel")
		r.Log.Panic(err)
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
		r.Log.Panic(err)
	}
}

func (r Rabbit) Boot() {
	once.Do(func() {
		var err error
		amqpURI = flag.String("amqp", os.Getenv("RABBIT_URL"), "AMQP URI")
		flag.Parse()
		conn, err = amqp.Dial(*amqpURI)

		if err != nil {
			err = errors.NewNotSupported(err, "error.rabbit-can-not-connect-to-server")
			r.Log.Panic(err)
		}

	})
}
