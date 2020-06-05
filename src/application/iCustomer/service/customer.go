package service

import (
	message "github.com/I-Reven/Hexagonal/src/domain/message/rabbit"
	queue "github.com/I-Reven/Hexagonal/src/framework/queue/rabbit"
	"github.com/I-Reven/Hexagonal/src/infrastructure/adapter/elasticsearch"
	repository "github.com/I-Reven/Hexagonal/src/infrastructure/repository/cassandra/customer"
)

type Customer struct {
	produce    queue.Produce
	repository repository.Customer
	syncIndex  elasticsearch.SyncIndex
}

func (s Customer) CreateProducer(customerName string) error {
	msg := message.CreateCustomer{CustomerName: customerName}

	return s.produce.ProduceMessage(msg)
}

func (s Customer) Create(customerName string) error {
	return s.repository.Migrate(customerName)
}

func (s Customer) SyncIndex(customerName string) error {
	return s.syncIndex.Send(customerName, "rooms")
}
