package job

import (
	"encoding/json"
	"github.com/I-Reven/Hexagonal/src/application/iRoomConsumer/service"
	"github.com/I-Reven/Hexagonal/src/domain/job"
	"github.com/I-Reven/Hexagonal/src/domain/message/rabbit"
	message "github.com/I-Reven/Hexagonal/src/domain/message/slack"
	"github.com/I-Reven/Hexagonal/src/framework/logger"
	"github.com/I-Reven/Hexagonal/src/framework/notification/slack"
	"github.com/I-Reven/Hexagonal/src/infrastructure/repository/redis/webHook"
	"github.com/juju/errors"
)

type (
	AddMetaData struct {
		tries      int
		message    rabbit.AddMetaData
		log        logger.Log
		slack      slack.Slack
		service    service.RoomConsumer
		repository webHook.WebHook
	}
)

func (j AddMetaData) Init(b []byte) (error, job.Job) {
	return json.Unmarshal(b, &j.message), j
}

func (j AddMetaData) Handler() error {
	return j.service.AddMetaData(j.message.CustomerName, j.message.RoomId, j.message.Key, j.message.Kind, j.message.Value)
}

func (j AddMetaData) Done() {}

func (j AddMetaData) Failed(err error) {
	err = errors.NewNotSupported(err, "error.job-failed")

	j.slack.Send(&message.FailedJob{
		JobName: "AddMetaData",
		Message: "Add Meta Data Job Failed",
		Error:   err,
	})

	j.log.Warn(err)
}

func (j AddMetaData) GetConfig() job.Config {
	return job.Config{Tries: 3}
}
