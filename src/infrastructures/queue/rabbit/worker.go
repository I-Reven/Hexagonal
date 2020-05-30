package rabbit

import (
	"github.com/I-Reven/Hexagonal/src/domains/job"
	"github.com/I-Reven/Hexagonal/src/domains/message/rabbit"
	"github.com/I-Reven/Hexagonal/src/infrastructures/logger"
	"github.com/juju/errors"
	"github.com/streadway/amqp"
)

type Worker struct {
	Log      logger.Log
	Consumer Consume
}

//AddWorker Add Worker for queue
func (w Worker) AddWorker(message rabbit.Message, jobs ...job.Job) {
	replies, err := w.Consumer.Message(message)

	if err != nil {
		err = errors.NewNotAssigned(err, "error.rabbit-can-not-consume-message")
		w.Log.Panic(err)
	}

	for r := range replies {
		if w.isValidMessage(message, r) {
			w.execute(r.Body, jobs...)
		}
	}
}

func (w Worker) isValidMessage(message rabbit.Message, r amqp.Delivery) bool {
	return r.Exchange == message.GetConsumerConfig().SourceExchange && r.RoutingKey == message.GetConsumerConfig().BindingKey
}

func (w Worker) execute(body []byte, jobs ...job.Job) {
	for i := len(jobs) - 1; i >= 0; i-- {
		w.try(body, jobs[i], 1)
	}
}

func (w Worker) try(body []byte, job job.Job, tries int) {
	var err error
	err, job = job.Init(body)
	if err == nil {
		err = job.Handler()
	}
	w.catch(body, job, tries, err)
}

func (w Worker) catch(body []byte, job job.Job, tries int, err error) {
	if err != nil {
		if job.GetConfig().Tries <= tries {
			w.try(body, job, tries+1)
		} else {
			job.Failed(err)
		}
	}
}
