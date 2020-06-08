package rabbit

type (
	AddUser struct {
		CustomerName string `json:"customer_name"`
		RoomId       int64  `json:"room_id"`
		UserId       int64  `json:"user_id"`
	}
)

func (AddUser) GetExchangeConfig() ExchangeConfig {
	return ExchangeConfig{
		Name:        "addUser-exchange",
		Type:        "direct",
		Durable:     true,
		AutoDeleted: false,
		Internal:    false,
		NoWait:      false,
		Arguments:   nil,
	}
}

func (AddUser) GetProducerConfig() ProducerConfig {
	return ProducerConfig{
		Exchange:   "addUser-exchange",
		RoutingKey: "addUser-key",
		Mandatory:  false,
		Immediate:  false,
	}
}

func (AddUser) GetConsumerConfig() ConsumerConfig {
	return ConsumerConfig{
		Name:           "addUser-queue",
		BindingKey:     "addUser-key",
		SourceExchange: "addUser-exchange",
		Consumer:       "amqp-addUser",
		AutoAck:        true,
		NoLocal:        false,
		Durable:        false,
		DeleteWhenUsed: true,
		Exclusive:      false,
	}
}
