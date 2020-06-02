package core

import (
	"github.com/I-Reven/Hexagonal/src/domain/entity"
	"gopkg.in/mgo.v2/bson"
)

type IAmAlive struct {
	entity.IAmAlive `bson:",inline"`
	core            Core
}

func (iAmAlive *IAmAlive) HttpTestSuccess() error {
	iAmAlive.SetContent("I Am Alive")
	iAmAlive.SetHttpSuccess(true)
	return iAmAlive.Save()
}

func (iAmAlive *IAmAlive) DbTestSuccess() error {
	iAmAlive.SetDbSuccess(true)
	return iAmAlive.Save()
}

func (iAmAlive *IAmAlive) ProducerTestSuccess() error {
	iAmAlive.SetProducerSuccess(true)
	return iAmAlive.Save()
}

func (iAmAlive *IAmAlive) ConsumerTestSuccess() error {
	iAmAlive.SetConsumerSuccess(true)
	return iAmAlive.Save()
}

func (iAmAlive *IAmAlive) CashTestSuccess() error {
	iAmAlive.SetCashSuccess(true)
	return iAmAlive.Save()
}

func (iAmAlive *IAmAlive) GrpcTestSuccess() error {
	iAmAlive.SetGrpcSuccess(true)
	return iAmAlive.Save()
}

func (iAmAlive *IAmAlive) Save() error {
	return iAmAlive.core.Connection().Collection("iAmAlive").Save(iAmAlive)
}

func (iAmAlive *IAmAlive) GetById(Id bson.ObjectId) error {
	return iAmAlive.core.Connection().Collection("iAmAlive").FindById(Id, iAmAlive)
}

func (iAmAlive *IAmAlive) GetLast() error {
	return iAmAlive.core.Connection().Collection("iAmAlive").Find(nil).Query.Sort("-created").One(iAmAlive)
}
