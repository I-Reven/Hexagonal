package core

import (
	"github.com/I-Reven/Hexagonal/src/applications/core/job"
	message "github.com/I-Reven/Hexagonal/src/domain/message/rabbit"
	"github.com/I-Reven/Hexagonal/src/framework/queue/rabbit"
)

type Worker struct {
	Worker rabbit.Worker
}

func (w Worker) Work() {
	go w.Worker.AddWorker(message.IAmAlive{}, job.IAmAliveJob{})
	go w.Worker.AddWorker(message.TrackRequest{}, job.RequestTracker{})
}
