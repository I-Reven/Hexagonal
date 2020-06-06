package job

import (
	"encoding/json"
	"github.com/I-Reven/Hexagonal/src/application/iCustomer/service"
	"github.com/I-Reven/Hexagonal/src/domain/job"
	"github.com/I-Reven/Hexagonal/src/domain/message/rabbit"
	message "github.com/I-Reven/Hexagonal/src/domain/message/slack"
	"github.com/I-Reven/Hexagonal/src/framework/logger"
	"github.com/I-Reven/Hexagonal/src/framework/notification/slack"
	"github.com/I-Reven/Hexagonal/src/infrastructure/repository/redis/webHook"
	"github.com/juju/errors"
	"os"
)

type (
	CreateCustomer struct {
		tries      int
		message    rabbit.CreateCustomer
		log        logger.Log
		slack      slack.Slack
		service    service.Customer
		repository webHook.WebHook
	}
)

func (j CreateCustomer) Init(b []byte) (error, job.Job) {
	return json.Unmarshal(b, &j.message), j
}

func (j CreateCustomer) Handler() error {
	err := j.service.Create(j.message.CustomerName)

	if err != nil {
		err := errors.NewNotSupported(err, "error.can-cot-migrate-customer-model")
		j.log.Error(err)
	}

	if err := j.service.SyncIndex(j.message.CustomerName); err != nil {
		err := errors.NewNotSupported(err, "error.can-not-sync-customer-index")
		j.log.Error(err)
		return err
	}

	return err
}

func (j CreateCustomer) Done() {
	j.slack.Send(&message.SuccessJob{
		JobName: "CreateCustomer",
		Message: "Create Customer " + j.message.CustomerName + " is Done",
	})
}

func (j CreateCustomer) Failed(err error) {
	err = errors.NewNotSupported(err, "error.job-failed")
	retryUrl, cancelUrl := j.getWebHooks()

	j.slack.Send(&message.FailedJob{
		JobName:   "CreateCustomer",
		Message:   "Create Customer Job Failed For Customer: " + j.message.CustomerName,
		RetryUrl:  retryUrl,
		CancelUrl: cancelUrl,
		Error:     err,
	})

	j.log.Warn(err)
}

func (j CreateCustomer) getWebHooks() (string, string) {
	retryUrl := ""
	cancelUrl := ""
	if data, er := json.Marshal(j.message); er == nil {
		if key, e := j.repository.Create("CreateCustomer", data); e == nil {
			retryUrl = os.Getenv("APP_URL") + ":80/web-hook/create-customer/" + key
			cancelUrl = os.Getenv("APP_URL") + ":80/web-hook/create-customer-cancel/" + key
		}
	}

	return retryUrl, cancelUrl
}

func (j CreateCustomer) GetConfig() job.Config {
	return job.Config{Tries: 1}
}
