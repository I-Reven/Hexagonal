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

// Test All Services
func Test() {
	iAmAlive := getEntity(testHttp())
	testCache(&iAmAlive)
	testProducer(&iAmAlive)
}

//Test Http
func testHttp() bson.ObjectId {
	iAmAlive := repository.IAmAlive{}
	err := iAmAlive.HttpTestSuccess()

	if err != nil {
		err = errors.NewNotSupported(err, "error.service-can-not-test-database")
		logger.Info(err)
	}

	return iAmAlive.GetId()
}

//Test Producer
func testProducer(iAmAlive *repository.IAmAlive) {
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

//Test Cache
func testCache(iAmAlive *repository.IAmAlive) {
	Cache := cache.Cache()
	key := "iAmAlive:" + string(iAmAlive.GetId())
	_, err := Cache.Set(key, iAmAlive.GetContent(), 10*time.Minute)
	_, err = Cache.Get(key)

	if err != nil {
		err = errors.NewNotSupported(err, "error.service-can-not-test-cache")
		logger.Info(err)
	} else {
		_ = iAmAlive.CashTestSuccess()
	}
}

//Get Entity
func getEntity(id bson.ObjectId) repository.IAmAlive {
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

func GetLastTest() (error, repository.IAmAlive) {
	iAmAlive := repository.IAmAlive{}
	return iAmAlive.GetLast(), iAmAlive
}
