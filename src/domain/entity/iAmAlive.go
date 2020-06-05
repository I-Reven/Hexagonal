package entity

import (
	"github.com/bxcodec/faker/v3"
	_ "github.com/bxcodec/faker/v3"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type (
	IAmAlive struct {
		Id              bson.ObjectId `bson:"_id,omitempty" json:"id" faker:"-"`
		HttpSuccess     bool          `bson:"httpSuccess" json:"httpSuccess"`
		ProducerSuccess bool          `bson:"producerSuccess" json:"producerSuccess"`
		DbSuccess       bool          `bson:"dbSuccess" json:"dbSuccess"`
		ConsumerSuccess bool          `bson:"consumerSuccess" json:"consumerSuccess"`
		CashSuccess     bool          `bson:"cashSuccess" json:"cashSuccess"`
		GrpcSuccess     bool          `bson:"grpcSuccess" json:"grpcSuccess"`
		Content         string        `bson:"content" json:"content" faker:"sentence"`
		Created         time.Time     `bson:"created" json:"created"`
		Modified        time.Time     `bson:"modified" json:"modified"`
	}
)

func (e *IAmAlive) GetId() bson.ObjectId           { return e.Id }
func (e *IAmAlive) SetId(id bson.ObjectId)         { e.Id = id }
func (e *IAmAlive) SetCreated(time time.Time)      { e.Created = time }
func (e *IAmAlive) SetModified(time time.Time)     { e.Modified = time }
func (e *IAmAlive) SetHttpSuccess(status bool)     { e.HttpSuccess = status }
func (e *IAmAlive) GetHttpSuccess() bool           { return e.HttpSuccess }
func (e *IAmAlive) SetProducerSuccess(status bool) { e.ProducerSuccess = status }
func (e *IAmAlive) GetProducerSuccess() bool       { return e.ProducerSuccess }
func (e *IAmAlive) SetDbSuccess(status bool)       { e.DbSuccess = status }
func (e *IAmAlive) GetDbSuccess() bool             { return e.DbSuccess }
func (e *IAmAlive) SetConsumerSuccess(status bool) { e.ConsumerSuccess = status }
func (e *IAmAlive) GetConsumerSuccess() bool       { return e.ConsumerSuccess }
func (e *IAmAlive) SetCashSuccess(status bool)     { e.CashSuccess = status }
func (e *IAmAlive) GetCashSuccess() bool           { return e.CashSuccess }
func (e *IAmAlive) SetGrpcSuccess(status bool)     { e.GrpcSuccess = status }
func (e *IAmAlive) GetGrpcSuccess() bool           { return e.GrpcSuccess }
func (e *IAmAlive) SetContent(content string)      { e.Content = content }
func (e *IAmAlive) GetContent() string             { return e.Content }
func (e *IAmAlive) getEntity() IAmAlive            { return *e }
func (e *IAmAlive) Factory() error                 { return faker.FakeData(e) }
