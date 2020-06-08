package rabbit

type (
	DeliverMessage struct {
		CustomerName string `json:"customer_name"`
		RoomId       int64  `json:"room_id"`
		MessageId    string `json:"message_id"`
		UserId       int64  `json:"user_id"`
	}
)

func (DeliverMessage) GetExchangeConfig() ExchangeConfig {
	return ExchangeConfig{
		Name:        "deliverMessage-exchange",
		Type:        "direct",
		Durable:     true,
		AutoDeleted: false,
		Internal:    false,
		NoWait:      false,
		Arguments:   nil,
	}
}

func (DeliverMessage) GetProducerConfig() ProducerConfig {
	return ProducerConfig{
		Exchange:   "deliverMessage-exchange",
		RoutingKey: "deliverMessage-key",
		Mandatory:  false,
		Immediate:  false,
	}
}

func (DeliverMessage) GetConsumerConfig() ConsumerConfig {
	return ConsumerConfig{
		Name:           "deliverMessage-queue",
		BindingKey:     "deliverMessage-key",
		SourceExchange: "deliverMessage-exchange",
		Consumer:       "amqp-deliverMessage",
		AutoAck:        true,
		NoLocal:        false,
		Durable:        false,
		DeleteWhenUsed: true,
		Exclusive:      false,
	}
}
