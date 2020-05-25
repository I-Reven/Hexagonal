package logger

import (
	"encoding/json"
	"github.com/I-Reven/Hexagonal/src/domains/entity"
	"github.com/I-Reven/Hexagonal/src/infrastructures/repository/cassandra/log"
	"github.com/juju/errors"
)

type (
	RequestLogger struct {
		entity *entity.Log
		log    log.Logger
	}
)

func Request() RequestLogger {
	request := RequestLogger{entity: &entity.Log{}, log: log.Log()}
	err := request.log.Create(request.entity)

	if err != nil {
		Error(errors.NewNotSupported(err, "Can not connect to cassandra logs"))
	}

	return request
}

func (r RequestLogger) Message(message string) {
	err := r.log.Update(r.entity.GetId(), message, r.entity.GetError())

	if err != nil {
		Error(errors.NewNotSupported(err, "Can not update log message"))
	}
}

func (r RequestLogger) Error(error error) {
	err := r.log.Update(r.entity.GetId(), r.entity.GetMessage(), error.Error())

	if err != nil {
		Error(errors.NewNotSupported(err, "Can not update log error"))
	}
}

func (r RequestLogger) Data(info interface{}) {
	data, err := json.Marshal(info)

	if err != nil {
		Info(errors.NewNotSupported(err, "Can not marshal interface log data"))
	}

	err = r.log.AddData(r.entity.GetId(), string(data))

	if err != nil {
		Error(errors.NewNotSupported(err, "Can not update log data"))
	}
}
