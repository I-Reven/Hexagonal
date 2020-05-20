package job

import (
	"encoding/json"
	"fmt"
	"github.com/I-Reven/Hexagonal/domains/job"
	message "github.com/I-Reven/Hexagonal/domains/message/rabbit/core"
	"github.com/I-Reven/Hexagonal/infrastructures/logger"
)

type (
	IAmAliveJob struct {
		Tries   int
		Message message.IAmAlive
	}
)

func (i IAmAliveJob) Init(b []byte) (error, job.Job) {
	return json.Unmarshal(b, &i.Message), i
}

func (i IAmAliveJob) Handler() error {
	fmt.Println(i.Message.Content)
	return nil
}

func (IAmAliveJob) Failed(err error) {
	logger.LOG().FatalE(err)
}

func (i IAmAliveJob) GetConfig() job.Config {
	return job.Config{Tries: i.Tries}
}
