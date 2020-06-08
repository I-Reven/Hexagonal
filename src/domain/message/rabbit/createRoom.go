package rabbit

type (
	CreateRoom struct {
		CustomerName string `json:"customer_name"`
		RoomId       int64  `json:"room_id"`
		UserId       int64  `json:"user_id"`
	}
)

func (CreateRoom) GetExchangeConfig() ExchangeConfig {
	return ExchangeConfig{
		Name:        "createRoom-exchange",
		Type:        "direct",
		Durable:     true,
		AutoDeleted: false,
		Internal:    false,
		NoWait:      false,
		Arguments:   nil,
	}
}

func (CreateRoom) GetProducerConfig() ProducerConfig {
	return ProducerConfig{
		Exchange:   "createRoom-exchange",
		RoutingKey: "createRoom-key",
		Mandatory:  false,
		Immediate:  false,
	}
}

func (CreateRoom) GetConsumerConfig() ConsumerConfig {
	return ConsumerConfig{
		Name:           "createRoom-queue",
		BindingKey:     "createRoom-key",
		SourceExchange: "createRoom-exchange",
		Consumer:       "amqp-createRoom",
		AutoAck:        true,
		NoLocal:        false,
		Durable:        false,
		DeleteWhenUsed: true,
		Exclusive:      false,
	}
}
