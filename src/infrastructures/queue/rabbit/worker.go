package rabbit

import (
	"github.com/I-Reven/Hexagonal/src/domains/job"
	"github.com/I-Reven/Hexagonal/src/domains/message/rabbit"
	"github.com/I-Reven/Hexagonal/src/infrastructures/logger"
	"github.com/juju/errors"
	"github.com/streadway/amqp"
)

//AddWorker Add Worker for queue
func AddWorker(message rabbit.Message, jobs ...job.Job) {
	replies, err := ConsumeMessage(message)

	if err != nil {
		err = errors.NewNotAssigned(err, "error.rabbit-can-not-consume-message")
		logger.Panic(err)
	}

	for r := range replies {
		if isValidMessage(message, r) {
			execute(r.Body, jobs...)
		}
	}
}

func isValidMessage(message rabbit.Message, r amqp.Delivery) bool {
	return r.Exchange == message.GetConsumerConfig().SourceExchange && r.RoutingKey == message.GetConsumerConfig().BindingKey
}

func execute(body []byte, jobs ...job.Job) {
	for i := len(jobs) - 1; i >= 0; i-- {
		try(body, jobs[i], 1)
	}
}

func try(body []byte, job job.Job, tries int) {
	var err error
	err, job = job.Init(body)
	if err == nil {
		err = job.Handler()
	}
	catch(body, job, tries, err)
}

func catch(body []byte, job job.Job, tries int, err error) {
	if err != nil {
		if job.GetConfig().Tries <= tries {
			try(body, job, tries+1)
		} else {
			job.Failed(err)
		}
	}
}
