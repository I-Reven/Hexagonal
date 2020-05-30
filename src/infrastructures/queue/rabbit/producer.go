package rabbit

import (
	"encoding/json"
	message "github.com/I-Reven/Hexagonal/src/domains/message/rabbit"
	"github.com/I-Reven/Hexagonal/src/infrastructures/logger"
	"github.com/juju/errors"
	"github.com/streadway/amqp"
	"time"
)

type Produce struct {
	Log    logger.Log
	Rabbit Rabbit
}

func (p Produce) ProduceMessage(message message.Message) error {
	p.Rabbit.Init(message)

	payload, err := json.Marshal(message)

	if err != nil {
		err = errors.NewNotSupported(err, "error.rabbit-can-not-produce-message")
		p.Log.Error(err)
		return err
	}

	err = ch.Publish(
		message.GetProducerConfig().Exchange,   // exchange
		message.GetProducerConfig().RoutingKey, // routing key
		message.GetProducerConfig().Mandatory,  // mandatory
		message.GetProducerConfig().Immediate,  // immediate
		amqp.Publishing{
			DeliveryMode: amqp.Transient,
			ContentType:  "application/json",
			Body:         payload,
			Timestamp:    time.Now(),
		})

	if err != nil {
		err = errors.NewNotSupported(err, "error.rabbit-can-not-publish-message-on-channel")
		p.Log.Error(err)
	}

	return err
}
