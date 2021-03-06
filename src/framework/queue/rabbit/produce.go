package rabbit

import (
	"encoding/json"
	message "github.com/I-Reven/Hexagonal/src/domain/message/rabbit"
	"github.com/I-Reven/Hexagonal/src/framework/logger"
	"github.com/juju/errors"
	"github.com/streadway/amqp"
	"time"
)

type Produce struct {
	log    logger.Log
	rabbit Rabbit
}

func (p *Produce) ProduceMessage(message message.Message) error {
	p.rabbit.Init(message)

	payload, err := json.Marshal(message)

	if err != nil {
		err = errors.NewNotSupported(err, "error.rabbit-can-not-produce-message")
		p.log.Error(err)
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
		p.log.Error(err)
	}

	return err
}
