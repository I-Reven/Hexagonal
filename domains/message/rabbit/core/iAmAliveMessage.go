package core

import (
	"github.com/I-Reven/Hexagonal/domains/message/rabbit"
)

type (
	IAmAlive struct {
		Id      string `json:"id"`
		Content string `json:"content"`
	}
)

func (IAmAlive) GetExchangeConfig() rabbit.ExchangeConfig {
	return rabbit.ExchangeConfig{
		Name:        "iAmAlive-exchange",
		Type:        "direct",
		Durable:     true,
		AutoDeleted: false,
		Internal:    false,
		NoWait:      false,
		Arguments:   nil,
	}
}

func (IAmAlive) GetProducerConfig() rabbit.ProducerConfig {
	return rabbit.ProducerConfig{
		Exchange:   "iAmAlive-exchange",
		RoutingKey: "iAmAlive-key",
		Mandatory:  false,
		Immediate:  false,
	}
}

func (IAmAlive) GetConsumerConfig() rabbit.ConsumerConfig {
	return rabbit.ConsumerConfig{
		Name:           "iAmAlive-queue",
		BindingKey:     "iAmAlive-key",
		SourceExchange: "iAmAlive-exchange",
		Consumer:       "amqp-iAmAlive",
		AutoAck:        true,
		NoLocal:        false,
		Durable:        false,
		DeleteWhenUsed: false,
		Exclusive:      false,
	}
}
