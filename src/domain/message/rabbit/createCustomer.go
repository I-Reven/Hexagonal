package rabbit

type (
	CreateCustomer struct {
		CustomerName string `json:"customer_name"`
	}
)

func (CreateCustomer) GetExchangeConfig() ExchangeConfig {
	return ExchangeConfig{
		Name:        "createCustomer-exchange",
		Type:        "direct",
		Durable:     true,
		AutoDeleted: false,
		Internal:    false,
		NoWait:      false,
		Arguments:   nil,
	}
}

func (CreateCustomer) GetProducerConfig() ProducerConfig {
	return ProducerConfig{
		Exchange:   "createCustomer-exchange",
		RoutingKey: "createCustomer-key",
		Mandatory:  false,
		Immediate:  false,
	}
}

func (CreateCustomer) GetConsumerConfig() ConsumerConfig {
	return ConsumerConfig{
		Name:           "createCustomer-queue",
		BindingKey:     "createCustomer-key",
		SourceExchange: "createCustomer-exchange",
		Consumer:       "amqp-createCustomer",
		AutoAck:        true,
		NoLocal:        false,
		Durable:        false,
		DeleteWhenUsed: true,
		Exclusive:      false,
	}
}
