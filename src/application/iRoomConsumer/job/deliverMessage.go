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
	DeliverMessage struct {
		tries      int
		message    rabbit.DeliverMessage
		log        logger.Log
		slack      slack.Slack
		service    service.RoomConsumer
		repository webHook.WebHook
	}
)

func (j DeliverMessage) Init(b []byte) (error, job.Job) {
	return json.Unmarshal(b, &j.message), j
}

func (j DeliverMessage) Handler() error {
	return j.service.DeliverMessage(j.message.CustomerName, j.message.RoomId, j.message.MessageId, j.message.UserId)
}

func (j DeliverMessage) Done() {}

func (j DeliverMessage) Failed(err error) {
	err = errors.NewNotSupported(err, "error.job-failed")

	j.slack.Send(&message.FailedJob{
		JobName: "DeliverMessage",
		Message: "Deliver Message Job Failed",
		Error:   err,
	})

	j.log.Warn(err)
}

func (j DeliverMessage) GetConfig() job.Config {
	return job.Config{Tries: 3}
}
