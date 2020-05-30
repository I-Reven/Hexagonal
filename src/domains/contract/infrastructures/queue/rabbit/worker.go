package rabbit

import (
	"github.com/I-Reven/Hexagonal/src/domains/job"
	"github.com/I-Reven/Hexagonal/src/domains/message/rabbit"
)

type Worker interface {
	AddWorker(message rabbit.Message, jobs ...job.Job)
}
