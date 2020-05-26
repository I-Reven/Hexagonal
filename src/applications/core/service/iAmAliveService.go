package service

import (
	message "github.com/I-Reven/Hexagonal/src/domains/message/rabbit"
	"github.com/I-Reven/Hexagonal/src/infrastructures/logger"
	"github.com/I-Reven/Hexagonal/src/infrastructures/queue/rabbit"
	repository "github.com/I-Reven/Hexagonal/src/infrastructures/repository/mongo/core"
	"github.com/I-Reven/Hexagonal/src/infrastructures/repository/redis/cache"
	"github.com/juju/errors"
	"gopkg.in/mgo.v2/bson"
	"time"
)

func Test() {
	iAmAlive := getEntity(testHttp())
	testCache(&iAmAlive)
	testProducer(&iAmAlive)
	logger.Debug("iAmAliveService.Test", iAmAlive)
}

func testHttp() bson.ObjectId {
	iAmAlive := repository.IAmAlive{}
	err := iAmAlive.HttpTestSuccess()

	if err != nil {
		err = errors.NewNotSupported(err, "error.service-can-not-test-database")
		logger.Warn(err)
	}

	logger.Debug("iAmAliveService.testHttp", iAmAlive)
	return iAmAlive.GetId()
}

func testProducer(iAmAlive *repository.IAmAlive) {
	mes := message.IAmAlive{
		Id:      iAmAlive.GetId(),
		Content: iAmAlive.GetContent(),
	}

	err := rabbit.ProduceMessage(mes)

	if err != nil {
		err = errors.NewNotSupported(err, "error.service-can-not-test-producer")
		logger.Warn(err)
	} else {
		_ = iAmAlive.ProducerTestSuccess()
	}

	logger.Debug("iAmAliveService.testProducer", mes)
}

func testCache(iAmAlive *repository.IAmAlive) {
	Cache := cache.Cache()
	key := "iAmAlive:" + string(iAmAlive.GetId())
	_, err := Cache.Set(key, iAmAlive.GetContent(), 10*time.Minute)
	_, err = Cache.Get(key)

	if err != nil {
		err = errors.NewNotSupported(err, "error.service-can-not-test-cache")
		logger.Warn(err)
	} else {
		_ = iAmAlive.CashTestSuccess()
	}

	logger.Debug("iAmAliveService.testCache", key, iAmAlive)
}

func getEntity(id bson.ObjectId) repository.IAmAlive {
	iAmAlive := repository.IAmAlive{}
	err := iAmAlive.GetById(id)

	if err != nil {
		err = errors.NewNotSupported(err, "error.service-can-not-get-data")
		logger.Warn(err)
	} else {
		_ = iAmAlive.DbTestSuccess()
	}

	logger.Debug("iAmAliveService.getEntity", id, iAmAlive)
	return iAmAlive
}

func GetLastTest() (repository.IAmAlive, error) {
	iAmAlive := repository.IAmAlive{}
	err := iAmAlive.GetLast()
	logger.Debug("iAmAliveService.GetLastTest", iAmAlive)
	return iAmAlive, err
}
