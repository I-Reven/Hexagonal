package core

import (
	"github.com/I-Reven/Hexagonal/applications/core/job"
	message "github.com/I-Reven/Hexagonal/domains/message/rabbit/core"
	"github.com/I-Reven/Hexagonal/infrastructures/queue/rabbit"
)

func Worker() {
	go rabbit.AddWorker(message.IAmAlive{}, job.IAmAliveJob{})
}
