package rabbit

type (
	AddMetaData struct {
		CustomerName string `json:"customer_name"`
		RoomId       int64  `json:"room_id"`
		Key          string `json:"key"`
		Kind         int64  `json:"kind"`
		Value        string `json:"value"`
	}
)

func (AddMetaData) GetExchangeConfig() ExchangeConfig {
	return ExchangeConfig{
		Name:        "addMetaData-exchange",
		Type:        "direct",
		Durable:     true,
		AutoDeleted: false,
		Internal:    false,
		NoWait:      false,
		Arguments:   nil,
	}
}

func (AddMetaData) GetProducerConfig() ProducerConfig {
	return ProducerConfig{
		Exchange:   "addMetaData-exchange",
		RoutingKey: "addMetaData-key",
		Mandatory:  false,
		Immediate:  false,
	}
}

func (AddMetaData) GetConsumerConfig() ConsumerConfig {
	return ConsumerConfig{
		Name:           "addMetaData-queue",
		BindingKey:     "addMetaData-key",
		SourceExchange: "addMetaData-exchange",
		Consumer:       "amqp-addMetaData",
		AutoAck:        true,
		NoLocal:        false,
		Durable:        false,
		DeleteWhenUsed: true,
		Exclusive:      false,
	}
}
