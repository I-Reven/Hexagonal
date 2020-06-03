package core

import (
	"github.com/I-Reven/Hexagonal/src/application/core/job"
	message "github.com/I-Reven/Hexagonal/src/domain/message/rabbit"
	"github.com/I-Reven/Hexagonal/src/framework/queue/rabbit"
)

type Worker struct {
	worker              rabbit.Worker
	iAmAliveMessage     message.IAmAlive
	iAmaLiveJob         job.IAmAliveJob
	trackRequestMessage message.TrackRequest
	trackRequestJob     job.TrackRequest
}

func (w Worker) Work() {
	go w.worker.AddWorker(w.iAmAliveMessage, w.iAmaLiveJob)
	go w.worker.AddWorker(w.trackRequestMessage, w.trackRequestJob)
}
