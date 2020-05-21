package service

import (
	message "github.com/I-Reven/Hexagonal/src/domains/message/rabbit"
	"github.com/I-Reven/Hexagonal/src/infrastructures/logger"
	"github.com/I-Reven/Hexagonal/src/infrastructures/queue/rabbit"
	repository "github.com/I-Reven/Hexagonal/src/infrastructures/repository/mongo/core"
	"github.com/juju/errors"
)

func Test() error {
	err, id := TestHttp()

	if err != nil {
		err = errors.NewNotSupported(err, "error.service-can-not-test-database")
		logger.Info(err)
		return err
	}

	iAmAlive := GetEntity(id)
	_ = iAmAlive.DbTestSuccess()

	err = TestProducer(iAmAlive)

	if err != nil {
		err = errors.NewNotSupported(err, "error.service-can-not-test-producer")
		logger.Info(err)
	} else {
		_ = iAmAlive.ProducerTestSuccess()
	}

	return err
}

func TestHttp() (error, string) {
	iAmAlive := repository.IAmAlive{}
	return iAmAlive.HttpTestSuccess(), string(iAmAlive.GetId())
}

func TestProducer(iAmAlive repository.IAmAlive) error {
	mes := message.IAmAlive{
		Id:      iAmAlive.GetId(),
		Content: iAmAlive.GetContent(),
	}

	return rabbit.ProduceMessage(mes)
}

func GetEntity(id string) repository.IAmAlive {
	iAmAlive := repository.IAmAlive{}
	err := iAmAlive.GetById(id)
	if err != nil {
		err = errors.NewNotSupported(err, "error.service-can-not-get-data")
		logger.Info(err)
	}

	return iAmAlive
}
