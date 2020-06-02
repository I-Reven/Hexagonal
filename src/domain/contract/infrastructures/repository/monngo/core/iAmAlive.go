package core

import "gopkg.in/mgo.v2/bson"

type IAmAlive interface {
	HttpTestSuccess() error
	DbTestSuccess() error
	ProducerTestSuccess() error
	ConsumerTestSuccess() error
	CashTestSuccess() error
	GrpcTestSuccess() error
	Save() error
	GetById(Id bson.ObjectId) error
	GetLast() error
}
