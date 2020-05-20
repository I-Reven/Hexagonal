package service

import (
	rabbit2 "github.com/I-Reven/Hexagonal/src/domains/message/rabbit"
	"github.com/I-Reven/Hexagonal/src/infrastructures/logger"
	"github.com/I-Reven/Hexagonal/src/infrastructures/queue/rabbit"
	repository "github.com/I-Reven/Hexagonal/src/infrastructures/repository/mongo/core"
	"github.com/juju/errors"
)

func Test() error {
	err, iAmAlive := TestDatabase()

	if err != nil {
		err = errors.NewNotSupported(err, "error.service-can-not-test-database")
		logger.Info(err)
	}

	logger.DebugF(iAmAlive, nil)

	err = TestProducer(iAmAlive)

	if err != nil {
		err = errors.NewNotSupported(err, "error.service-can-not-test-producer")
		logger.Info(err)
	}

	return err
}

func TestDatabase() (error, repository.IAmAlive) {
	iAmAlive := repository.IAmAlive{}
	return iAmAlive.Add(), iAmAlive
}

func TestProducer(iAmAlive repository.IAmAlive) error {
	mes := rabbit2.IAmAlive{
		Id:      string(iAmAlive.GetId()),
		Content: iAmAlive.GetContent(),
	}

	return rabbit.ProduceMessage(mes)
}
