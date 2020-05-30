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
		Log     logger.Log
	}
)

func (j IAmAliveJob) Init(b []byte) (error, job.Job) {
	return json.Unmarshal(b, &j.Message), j
}

func (j IAmAliveJob) Handler() error {
	iAmAlive := repository.IAmAlive{}
	err := iAmAlive.GetById(j.Message.Id)

	if err != nil {
		return err
	}

	return iAmAlive.ConsumerTestSuccess()
}

func (j IAmAliveJob) Failed(err error) {
	err = errors.NewNotSupported(err, "error.job-failed")

	slack.Send(&message.FailedJob{
		JobName: "IAmAliveJob",
		Message: "I Am Alive Job Failed",
		Error:   err,
	})

	j.Log.Warn(err)
}

func (j IAmAliveJob) GetConfig() job.Config {
	return job.Config{Tries: j.Tries}
}
