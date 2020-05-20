package rabbit

import (
	"flag"
	"github.com/I-Reven/Hexagonal/domains/message/rabbit"
	"github.com/I-Reven/Hexagonal/infrastructures/logger"
	"github.com/streadway/amqp"
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
		amqpURI = flag.String("amqp", "amqp://kousha:kousha@localhost:5672/", "AMQP URI")
		flag.Parse()
		conn, err = amqp.Dial(*amqpURI)
		logger.LOG().FatalE(err)
	})
}

func Init(message rabbit.Message) {
	Boot()
	var err error
	ch, err = conn.Channel()
	logger.LOG().FatalE(err)
	err = ch.ExchangeDeclare(
		message.GetExchangeConfig().Name,        // name
		message.GetExchangeConfig().Type,        // type
		message.GetExchangeConfig().Durable,     // durable
		message.GetExchangeConfig().AutoDeleted, // auto-deleted
		message.GetExchangeConfig().Internal,    // internal
		message.GetExchangeConfig().NoWait,      // noWait
		message.GetExchangeConfig().Arguments,   // arguments
	)
	logger.LOG().FatalE(err)
}
