package job

import (
	"encoding/json"
	"fmt"
	"github.com/I-Reven/Hexagonal/src/domains/job"
	"github.com/I-Reven/Hexagonal/src/domains/message/rabbit"
	"github.com/I-Reven/Hexagonal/src/infrastructures/logger"
	"github.com/juju/errors"
)

type (
	IAmAliveJob struct {
		Tries   int
		Message rabbit.IAmAlive
	}
)

func (i IAmAliveJob) Init(b []byte) (error, job.Job) {
	return json.Unmarshal(b, &i.Message), i
}

func (i IAmAliveJob) Handler() error {
	fmt.Println(i.Message.Content)
	return nil
}

func (IAmAliveJob) Failed(err error) {
	err = errors.NewNotSupported(err, "error.job-failed")
	logger.Info(err)
}

func (i IAmAliveJob) GetConfig() job.Config {
	return job.Config{Tries: i.Tries}
}
