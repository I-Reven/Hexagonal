package rabbit

type (
	SeenMessage struct {
		CustomerName string `json:"customer_name"`
		RoomId       int64  `json:"room_id"`
		MessageId    string `json:"message_id"`
		UserId       int64  `json:"user_id"`
	}
)

func (SeenMessage) GetExchangeConfig() ExchangeConfig {
	return ExchangeConfig{
		Name:        "seenMessage-exchange",
		Type:        "direct",
		Durable:     true,
		AutoDeleted: false,
		Internal:    false,
		NoWait:      false,
		Arguments:   nil,
	}
}

func (SeenMessage) GetProducerConfig() ProducerConfig {
	return ProducerConfig{
		Exchange:   "seenMessage-exchange",
		RoutingKey: "seenMessage-key",
		Mandatory:  false,
		Immediate:  false,
	}
}

func (SeenMessage) GetConsumerConfig() ConsumerConfig {
	return ConsumerConfig{
		Name:           "seenMessage-queue",
		BindingKey:     "seenMessage-key",
		SourceExchange: "seenMessage-exchange",
		Consumer:       "amqp-seenMessage",
		AutoAck:        true,
		NoLocal:        false,
		Durable:        false,
		DeleteWhenUsed: true,
		Exclusive:      false,
	}
}
