package job

import (
	"encoding/json"
	"github.com/I-Reven/Hexagonal/src/domains/job"
	"github.com/I-Reven/Hexagonal/src/domains/message/rabbit"
	message "github.com/I-Reven/Hexagonal/src/domains/message/slack"
	"github.com/I-Reven/Hexagonal/src/infrastructures/logger"
	"github.com/I-Reven/Hexagonal/src/infrastructures/notification/slack"
	repository "github.com/I-Reven/Hexagonal/src/infrastructures/repository/mongo/core"
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
	iAmAlive := repository.IAmAlive{}
	err := iAmAlive.GetById(i.Message.Id)

	if err != nil {
		return err
	}

	return iAmAlive.ConsumerTestSuccess()
}

func (IAmAliveJob) Failed(err error) {
	err = errors.NewNotSupported(err, "error.job-failed")

	slack.Send(&message.FailedJob{
		JobName: "IAmAliveJob",
		Message: "I Am Alive Job Failed",
		Error:   err,
	})

	logger.Warn(err)
}

func (i IAmAliveJob) GetConfig() job.Config {
	return job.Config{Tries: i.Tries}
}
