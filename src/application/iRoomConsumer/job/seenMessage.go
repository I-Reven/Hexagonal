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
	SeenMessage struct {
		tries      int
		message    rabbit.SeenMessage
		log        logger.Log
		slack      slack.Slack
		service    service.RoomConsumer
		repository webHook.WebHook
	}
)

func (j SeenMessage) Init(b []byte) (error, job.Job) {
	return json.Unmarshal(b, &j.message), j
}

func (j SeenMessage) Handler() error {
	return j.service.SeenMessage(j.message.CustomerName, j.message.RoomId, j.message.MessageId, j.message.UserId)
}

func (j SeenMessage) Done() {}

func (j SeenMessage) Failed(err error) {
	err = errors.NewNotSupported(err, "error.job-failed")

	j.slack.Send(&message.FailedJob{
		JobName: "SeenMessage",
		Message: "Seen Message Job Failed",
		Error:   err,
	})

	j.log.Warn(err)
}

func (j SeenMessage) GetConfig() job.Config {
	return job.Config{Tries: 3}
}
