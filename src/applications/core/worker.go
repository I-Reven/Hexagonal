package core

import (
	"github.com/I-Reven/Hexagonal/src/applications/core/job"
	message "github.com/I-Reven/Hexagonal/src/domains/message/rabbit"
	"github.com/I-Reven/Hexagonal/src/infrastructures/queue/rabbit"
)

//Worker Job
func worker() {
	go rabbit.AddWorker(message.IAmAlive{}, job.IAmAliveJob{})
}
