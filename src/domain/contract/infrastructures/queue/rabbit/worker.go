package rabbit

import (
	"github.com/I-Reven/Hexagonal/src/domain/job"
	"github.com/I-Reven/Hexagonal/src/domain/message/rabbit"
)

type Worker interface {
	AddWorker(message rabbit.Message, jobs ...job.Job)
}
