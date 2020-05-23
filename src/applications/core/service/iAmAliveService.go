package service

import (
	message "github.com/I-Reven/Hexagonal/src/domains/message/rabbit"
	"github.com/I-Reven/Hexagonal/src/infrastructures/logger"
	"github.com/I-Reven/Hexagonal/src/infrastructures/queue/rabbit"
	repository "github.com/I-Reven/Hexagonal/src/infrastructures/repository/mongo/core"
	"github.com/juju/errors"
	"gopkg.in/mgo.v2/bson"
)

func Test() {
	id := TestHttp()
	iAmAlive := GetEntity(id)
	TestProducer(iAmAlive)
}

func TestHttp() bson.ObjectId {
	iAmAlive := repository.IAmAlive{}
	err := iAmAlive.HttpTestSuccess()

	if err != nil {
		err = errors.NewNotSupported(err, "error.service-can-not-test-database")
		logger.Info(err)
	}

	return iAmAlive.GetId()
}

func TestProducer(iAmAlive repository.IAmAlive) {
	mes := message.IAmAlive{
		Id:      iAmAlive.GetId(),
		Content: iAmAlive.GetContent(),
	}

	err := rabbit.ProduceMessage(mes)

	if err != nil {
		err = errors.NewNotSupported(err, "error.service-can-not-test-producer")
		logger.Info(err)
	} else {
		_ = iAmAlive.ProducerTestSuccess()
	}
}

func GetEntity(id bson.ObjectId) repository.IAmAlive {
	iAmAlive := repository.IAmAlive{}
	err := iAmAlive.GetById(id)

	if err != nil {
		err = errors.NewNotSupported(err, "error.service-can-not-get-data")
		logger.Info(err)
	} else {
		_ = iAmAlive.DbTestSuccess()
	}

	return iAmAlive
}
