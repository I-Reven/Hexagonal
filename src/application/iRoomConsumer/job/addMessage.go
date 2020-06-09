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
	AddMessage struct {
		tries      int
		message    rabbit.AddMessage
		log        logger.Log
		slack      slack.Slack
		service    service.RoomConsumer
		repository webHook.WebHook
	}
)

func (j AddMessage) Init(b []byte) (error, job.Job) {
	return json.Unmarshal(b, &j.message), j
}

func (j AddMessage) Handler() error {
	return j.service.AddMessage(j.message.CustomerName, j.message.RoomId, j.message.UserId, j.message.MessageId, j.message.Content, j.message.Kind)
}

func (j AddMessage) Done() {}

func (j AddMessage) Failed(err error) {
	err = errors.NewNotSupported(err, "error.job-failed")

	j.slack.Send(&message.FailedJob{
		JobName: "AddMessage",
		Message: "Add Message Job Failed",
		Error:   err,
	})

	j.log.Warn(err)
}

func (j AddMessage) GetConfig() job.Config {
	return job.Config{Tries: 3}
}
