package job

import (
	"encoding/json"
	"github.com/I-Reven/Hexagonal/src/domain/job"
	"github.com/I-Reven/Hexagonal/src/domain/message/rabbit"
	message "github.com/I-Reven/Hexagonal/src/domain/message/slack"
	"github.com/I-Reven/Hexagonal/src/framework/logger"
	"github.com/I-Reven/Hexagonal/src/framework/notification/slack"
	repository "github.com/I-Reven/Hexagonal/src/infrastructure/repository/mongo/core"
	"github.com/juju/errors"
)

type (
	IAmAliveJob struct {
		tries   int
		message rabbit.IAmAlive
		log     logger.Log
		slack   slack.Slack
	}
)

func (j IAmAliveJob) Init(b []byte) (error, job.Job) {
	return json.Unmarshal(b, &j.message), j
}

func (j IAmAliveJob) Handler() error {
	iAmAlive := repository.IAmAlive{}
	err := iAmAlive.GetById(j.message.Id)

	if err != nil {
		return err
	}

	return iAmAlive.ConsumerTestSuccess()
}

func (j IAmAliveJob) Failed(err error) {
	err = errors.NewNotSupported(err, "error.job-failed")

	j.slack.Send(&message.FailedJob{
		JobName: "IAmAliveJob",
		Message: "I Am Alive Job Failed",
		Error:   err,
	})

	j.log.Warn(err)
}

func (j IAmAliveJob) GetConfig() job.Config {
	return job.Config{Tries: 1}
}
