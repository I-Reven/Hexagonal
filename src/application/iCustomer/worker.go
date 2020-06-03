package core

import (
	"github.com/I-Reven/Hexagonal/src/application/iCustomer/job"
	message "github.com/I-Reven/Hexagonal/src/domain/message/rabbit"
	queue "github.com/I-Reven/Hexagonal/src/framework/queue/rabbit"
)

type Worker struct {
	worker                queue.Worker
	createCustomerJob     job.CreateCustomer
	createCustomerMessage message.CreateCustomer
}

func (w Worker) Work() {
	go w.worker.AddWorker(w.createCustomerMessage, w.createCustomerJob)
}
