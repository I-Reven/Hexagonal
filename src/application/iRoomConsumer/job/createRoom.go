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
	CreateRoom struct {
		tries      int
		message    rabbit.CreateRoom
		log        logger.Log
		slack      slack.Slack
		service    service.RoomConsumer
		repository webHook.WebHook
	}
)

func (j CreateRoom) Init(b []byte) (error, job.Job) {
	return json.Unmarshal(b, &j.message), j
}

func (j CreateRoom) Handler() error {
	return j.service.Create(j.message.CustomerName, j.message.RoomId, j.message.UserId)
}

func (j CreateRoom) Done() {}

func (j CreateRoom) Failed(err error) {
	err = errors.NewNotSupported(err, "error.job-failed")

	j.slack.Send(&message.FailedJob{
		JobName: "CreateRoom",
		Message: "Create Room Job Failed",
		Error:   err,
	})

	j.log.Warn(err)
}

func (j CreateRoom) GetConfig() job.Config {
	return job.Config{Tries: 3}
}
