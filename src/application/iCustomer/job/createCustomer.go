package job

import (
	"encoding/json"
	"github.com/I-Reven/Hexagonal/src/application/iCustomer/service"
	"github.com/I-Reven/Hexagonal/src/domain/job"
	"github.com/I-Reven/Hexagonal/src/domain/message/rabbit"
	message "github.com/I-Reven/Hexagonal/src/domain/message/slack"
	"github.com/I-Reven/Hexagonal/src/framework/logger"
	"github.com/I-Reven/Hexagonal/src/framework/notification/slack"
	"github.com/juju/errors"
)

type (
	CreateCustomer struct {
		tries   int
		message rabbit.CreateCustomer
		log     logger.Log
		slack   slack.Slack
		service service.Customer
	}
)

func (j CreateCustomer) Init(b []byte) (error, job.Job) {
	return json.Unmarshal(b, &j.message), j
}

func (j CreateCustomer) Handler() error {
	j.log.TraceLn("get message create new customer " + j.message.CustomerName)
	if err := j.service.Create(j.message.CustomerName); err != nil {
		err := errors.NewNotSupported(err, "error.can-cot-migrate-customer-model")
		j.log.Error(err)
		return err
	}

	if err := j.service.SyncIndex(j.message.CustomerName); err != nil {
		err := errors.NewNotSupported(err, "error.can-not-sync-customer-index")
		j.log.Error(err)
		return err
	}

	return nil
}

func (j CreateCustomer) Failed(err error) {
	err = errors.NewNotSupported(err, "error.job-failed")

	j.slack.Send(&message.FailedJob{
		JobName: "CreateCustomer",
		Message: "Create Customer Job Failed For Partner " + j.message.CustomerName,
		Error:   err,
	})

	j.log.Warn(err)
}

func (j CreateCustomer) GetConfig() job.Config {
	return job.Config{Tries: 2}
}
