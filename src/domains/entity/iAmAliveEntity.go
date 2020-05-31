package entity

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type (
	IAmAlive struct {
		Id              bson.ObjectId `bson:"_id,omitempty" json:"id"`
		HttpSuccess     bool          `bson:"httpSuccess" json:"httpSuccess"`
		ProducerSuccess bool          `bson:"producerSuccess" json:"producerSuccess"`
		DbSuccess       bool          `bson:"dbSuccess" json:"dbSuccess"`
		ConsumerSuccess bool          `bson:"consumerSuccess" json:"consumerSuccess"`
		CashSuccess     bool          `bson:"cashSuccess" json:"cashSuccess"`
		GrpcSuccess     bool          `bson:"grpcSuccess" json:"grpcSuccess"`
		Content         string        `bson:"content" json:"content"`
		Created         time.Time     `bson:"created" json:"created"`
		Modified        time.Time     `bson:"modified" json:"modified"`
	}
)

func (iAmAlive *IAmAlive) GetId() bson.ObjectId       { return iAmAlive.Id }
func (iAmAlive *IAmAlive) SetId(id bson.ObjectId)     { iAmAlive.Id = id }
func (iAmAlive *IAmAlive) SetCreated(time time.Time)  { iAmAlive.Created = time }
func (iAmAlive *IAmAlive) SetModified(time time.Time) { iAmAlive.Modified = time }

func (iAmAlive *IAmAlive) SetHttpSuccess(status bool)     { iAmAlive.HttpSuccess = status }
func (iAmAlive *IAmAlive) GetHttpSuccess() bool           { return iAmAlive.HttpSuccess }
func (iAmAlive *IAmAlive) SetProducerSuccess(status bool) { iAmAlive.ProducerSuccess = status }
func (iAmAlive *IAmAlive) GetProducerSuccess() bool       { return iAmAlive.ProducerSuccess }
func (iAmAlive *IAmAlive) SetDbSuccess(status bool)       { iAmAlive.DbSuccess = status }
func (iAmAlive *IAmAlive) GetDbSuccess() bool             { return iAmAlive.DbSuccess }
func (iAmAlive *IAmAlive) SetConsumerSuccess(status bool) { iAmAlive.ConsumerSuccess = status }
func (iAmAlive *IAmAlive) GetConsumerSuccess() bool       { return iAmAlive.ConsumerSuccess }
func (iAmAlive *IAmAlive) SetCashSuccess(status bool)     { iAmAlive.CashSuccess = status }
func (iAmAlive *IAmAlive) GetCashSuccess() bool           { return iAmAlive.CashSuccess }
func (iAmAlive *IAmAlive) SetGrpcSuccess(status bool)     { iAmAlive.GrpcSuccess = status }
func (iAmAlive *IAmAlive) GetGrpcSuccess() bool           { return iAmAlive.GrpcSuccess }
func (iAmAlive *IAmAlive) SetContent(content string)      { iAmAlive.Content = content }
func (iAmAlive *IAmAlive) GetContent() string             { return iAmAlive.Content }
func (iAmAlive *IAmAlive) getEntity() IAmAlive            { return *iAmAlive }
