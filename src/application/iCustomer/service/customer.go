package service

import (
	message "github.com/I-Reven/Hexagonal/src/domain/message/rabbit"
	queue "github.com/I-Reven/Hexagonal/src/framework/queue/rabbit"
	repository "github.com/I-Reven/Hexagonal/src/infrastructure/repository/cassandra/customer"
)

type Customer struct {
	produce    queue.Produce
	repository repository.Customer
}

func (s Customer) CreateProducer(customerName string) error {
	msg := message.CreateCustomer{CustomerName: customerName}

	return s.produce.ProduceMessage(msg)
}

func (s Customer) Create(customerName string) error {
	return s.repository.Migrate(customerName)
}
