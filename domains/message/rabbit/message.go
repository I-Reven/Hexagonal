package rabbit

import "github.com/streadway/amqp"

type (
	ExchangeConfig struct {
		Name        string		// name of the exchange
		Type        string		// type
		Durable     bool		// durable
		AutoDeleted bool		// delete when complete
		Internal    bool		// internal
		NoWait      bool		// noWait
		Arguments   amqp.Table	// arguments
	}

	ProducerConfig struct {
		Exchange   string		// exchange
		RoutingKey string		// routing key
		Mandatory  bool			// mandatory
		Immediate  bool			// immediate
	}

	ConsumerConfig struct {
		Name string				// name, leave empty to generate a unique name
		BindingKey string		// binding key == routing key
		SourceExchange string	// source exchange == exchange
		Consumer string			// consumer
		AutoAck bool			// auto ack
		NoLocal bool			// no local
		Durable bool			// durable
		DeleteWhenUsed bool		// delete when used
		Exclusive bool			// exclusive
	}

	Message interface {
		GetExchangeConfig() ExchangeConfig
		GetProducerConfig() ProducerConfig
		GetConsumerConfig() ConsumerConfig
	}
)
