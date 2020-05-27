package rabbit

import "gopkg.in/mgo.v2/bson"

type (
	IAmAlive struct {
		Id      bson.ObjectId `json:"id"`
		Content string        `json:"content"`
	}
)

func (IAmAlive) GetExchangeConfig() ExchangeConfig {
	return ExchangeConfig{
		Name:        "iAmAlive-exchange",
		Type:        "direct",
		Durable:     true,
		AutoDeleted: false,
		Internal:    false,
		NoWait:      false,
		Arguments:   nil,
	}
}

func (IAmAlive) GetProducerConfig() ProducerConfig {
	return ProducerConfig{
		Exchange:   "iAmAlive-exchange",
		RoutingKey: "iAmAlive-key",
		Mandatory:  false,
		Immediate:  false,
	}
}

func (IAmAlive) GetConsumerConfig() ConsumerConfig {
	return ConsumerConfig{
		Name:           "iAmAlive-queue",
		BindingKey:     "iAmAlive-key",
		SourceExchange: "iAmAlive-exchange",
		Consumer:       "amqp-iAmAlive",
		AutoAck:        true,
		NoLocal:        false,
		Durable:        false,
		DeleteWhenUsed: true,
		Exclusive:      false,
	}
}
