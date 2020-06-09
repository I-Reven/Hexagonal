package core

import (
	"github.com/I-Reven/Hexagonal/src/application/iRoomConsumer/job"
	message "github.com/I-Reven/Hexagonal/src/domain/message/rabbit"
	"github.com/I-Reven/Hexagonal/src/framework/queue/rabbit"
)

type Worker struct {
	worker                rabbit.Worker
	createRoomMessage     message.CreateRoom
	addUserMessage        message.AddUser
	addMessageMessage     message.AddMessage
	addMetaDataMessage    message.AddMetaData
	seenMessageMessage    message.SeenMessage
	deliverMessageMessage message.DeliverMessage
	createRoomJob         job.CreateRoom
	addUserJob            job.AddUser
	addMessageJob         job.AddMessage
	addMetaDataJob        job.AddMetaData
	seenMessageJob        job.SeenMessage
	deliverMessageJob     job.DeliverMessage
}

func (w Worker) Work() {
	go w.worker.AddWorker(w.createRoomMessage, w.createRoomJob)
	go w.worker.AddWorker(w.addUserMessage, w.addUserJob)
	go w.worker.AddWorker(w.addMessageMessage, w.addMessageJob)
	go w.worker.AddWorker(w.addMetaDataMessage, w.addMetaDataJob)
	go w.worker.AddWorker(w.seenMessageMessage, w.seenMessageJob)
	go w.worker.AddWorker(w.deliverMessageMessage, w.deliverMessageJob)
}
