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
	AddUser struct {
		tries      int
		message    rabbit.AddUser
		log        logger.Log
		slack      slack.Slack
		service    service.RoomConsumer
		repository webHook.WebHook
	}
)

func (j AddUser) Init(b []byte) (error, job.Job) {
	return json.Unmarshal(b, &j.message), j
}

func (j AddUser) Handler() error {
	return j.service.AddUser(j.message.CustomerName, j.message.RoomId, j.message.UserId)
}

func (j AddUser) Done() {}

func (j AddUser) Failed(err error) {
	err = errors.NewNotSupported(err, "error.job-failed")

	j.slack.Send(&message.FailedJob{
		JobName: "AddUser",
		Message: "Add User Job Failed",
		Error:   err,
	})

	j.log.Warn(err)
}

func (j AddUser) GetConfig() job.Config {
	return job.Config{Tries: 3}
}
