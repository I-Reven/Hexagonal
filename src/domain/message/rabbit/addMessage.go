package rabbit

type (
	AddMessage struct {
		CustomerName string `json:"customer_name"`
		RoomId       int64  `json:"room_id"`
		UserId       int64  `json:"user_id"`
		MessageId    string `json:"message_id"`
		Content      string `json:"content"`
		Kind         int64  `json:"kind"`
	}
)

func (AddMessage) GetExchangeConfig() ExchangeConfig {
	return ExchangeConfig{
		Name:        "addMessage-exchange",
		Type:        "direct",
		Durable:     true,
		AutoDeleted: false,
		Internal:    false,
		NoWait:      false,
		Arguments:   nil,
	}
}

func (AddMessage) GetProducerConfig() ProducerConfig {
	return ProducerConfig{
		Exchange:   "addMessage-exchange",
		RoutingKey: "addMessage-key",
		Mandatory:  false,
		Immediate:  false,
	}
}

func (AddMessage) GetConsumerConfig() ConsumerConfig {
	return ConsumerConfig{
		Name:           "addMessage-queue",
		BindingKey:     "addMessage-key",
		SourceExchange: "addMessage-exchange",
		Consumer:       "amqp-addMessage",
		AutoAck:        true,
		NoLocal:        false,
		Durable:        false,
		DeleteWhenUsed: true,
		Exclusive:      false,
	}
}
