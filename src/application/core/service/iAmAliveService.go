package service

import (
	domain "github.com/I-Reven/Hexagonal/src/domain/grpc"
	message "github.com/I-Reven/Hexagonal/src/domain/message/rabbit"
	"github.com/I-Reven/Hexagonal/src/framework/logger"
	"github.com/I-Reven/Hexagonal/src/framework/queue/rabbit"
	"github.com/I-Reven/Hexagonal/src/infrastructure/grpc/core/client"
	repository "github.com/I-Reven/Hexagonal/src/infrastructure/repository/mongo/core"
	"github.com/I-Reven/Hexagonal/src/infrastructure/repository/redis/cache"
	"github.com/juju/errors"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type IAamAliveService struct {
	log     logger.Log
	Produce rabbit.Produce
	cache   cache.Cache
	client  client.Ping
}

func (s IAamAliveService) Test() {
	iAmAlive := s.getEntity(s.testHttp())
	s.testGrpc(&iAmAlive)
	s.testCache(&iAmAlive)
	s.testProducer(&iAmAlive)

	s.log.Debug("iAmAliveService.Test", iAmAlive)
}

func (s IAamAliveService) testHttp() bson.ObjectId {
	iAmAlive := repository.IAmAlive{}
	err := iAmAlive.HttpTestSuccess()

	if err != nil {
		err = errors.NewNotSupported(err, "error.service-can-not-test-database")
		s.log.Warn(err)
	}

	return iAmAlive.GetId()
}

func (s IAamAliveService) testProducer(iAmAlive *repository.IAmAlive) {
	mes := message.IAmAlive{
		Id:      iAmAlive.GetId(),
		Content: iAmAlive.GetContent(),
	}

	err := s.Produce.ProduceMessage(mes)

	if err != nil {
		err = errors.NewNotSupported(err, "error.service-can-not-test-producer")
		s.log.Warn(err)
	} else {
		_ = iAmAlive.ProducerTestSuccess()
	}
}

func (s IAamAliveService) testCache(iAmAlive *repository.IAmAlive) {
	Cache := s.cache.Init()
	key := "iAmAlive:" + string(iAmAlive.GetId())
	err := Cache.Set(key, iAmAlive.GetContent(), 10*time.Minute)
	_, err = Cache.Get(key)

	if err != nil {
		err = errors.NewNotSupported(err, "error.service-can-not-test-cache")
		s.log.Warn(err)
	} else {
		_ = iAmAlive.CashTestSuccess()
	}
}

func (s IAamAliveService) testGrpc(iAmAlive *repository.IAmAlive) {
	_, err := s.client.Ping(":82", &domain.PingRequest{Message: "PING"})

	if err != nil {
		err = errors.NewNotSupported(err, "error.service-can-not-test-cache")
		s.log.Warn(err)
	} else {
		_ = iAmAlive.GrpcTestSuccess()
	}
}

func (s IAamAliveService) getEntity(id bson.ObjectId) repository.IAmAlive {
	iAmAlive := repository.IAmAlive{}
	err := iAmAlive.GetById(id)

	if err != nil {
		err = errors.NewNotSupported(err, "error.service-can-not-get-data")
		s.log.Warn(err)
	} else {
		_ = iAmAlive.DbTestSuccess()
	}

	return iAmAlive
}

func (s IAamAliveService) GetLastTest() (repository.IAmAlive, error) {
	iAmAlive := repository.IAmAlive{}
	err := iAmAlive.GetLast()
	return iAmAlive, err
}
