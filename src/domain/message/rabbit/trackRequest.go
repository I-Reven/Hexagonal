package rabbit

type (
	TrackRequest struct {
		Id string `json:"id"`
	}
)

func (TrackRequest) GetExchangeConfig() ExchangeConfig {
	return ExchangeConfig{
		Name:        "trackRequest-exchange",
		Type:        "direct",
		Durable:     true,
		AutoDeleted: false,
		Internal:    false,
		NoWait:      false,
		Arguments:   nil,
	}
}

func (TrackRequest) GetProducerConfig() ProducerConfig {
	return ProducerConfig{
		Exchange:   "trackRequest-exchange",
		RoutingKey: "trackRequest-key",
		Mandatory:  false,
		Immediate:  false,
	}
}

func (TrackRequest) GetConsumerConfig() ConsumerConfig {
	return ConsumerConfig{
		Name:           "trackRequest-queue",
		BindingKey:     "trackRequest-key",
		SourceExchange: "trackRequest-exchange",
		Consumer:       "amqp-trackRequest",
		AutoAck:        true,
		NoLocal:        false,
		Durable:        false,
		DeleteWhenUsed: true,
		Exclusive:      false,
	}
}
