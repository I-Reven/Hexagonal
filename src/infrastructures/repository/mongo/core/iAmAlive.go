package core

import (
	entity "github.com/I-Reven/Hexagonal/src/domains/entity/core"
	"gopkg.in/mgo.v2/bson"
)

type IAmAlive struct {
	entity.IAmAlive `bson:",inline"`
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

func (iAmAlive *IAmAlive) SocketTestSuccess() error {
	iAmAlive.SetCashSuccess(true)
	return iAmAlive.Save()
}

func (iAmAlive *IAmAlive) Save() error {
	return Mongo().Collection("iAmAlive").Save(iAmAlive)
}

func (iAmAlive *IAmAlive) GetById(Id bson.ObjectId) error {
	return Mongo().Collection("iAmAlive").FindById(Id, iAmAlive)
}
