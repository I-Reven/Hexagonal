package rabbit

import (
	"encoding/json"
	message "github.com/I-Reven/Hexagonal/domains/message/rabbit"
	"github.com/streadway/amqp"
	"time"
)

func ProduceMessage(message message.Message) error {
	Init(message)

	payload, err := json.Marshal(message)

	if err != nil {
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

	return err
}
