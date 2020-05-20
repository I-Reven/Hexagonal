package core

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type (
	IAmAlive struct {
		Id       bson.ObjectId `bson:"_id,omitempty" json:"id"`
		Content  string        `bson:"content" json:"content"`
		Created  time.Time     `bson:"created" json:"created"`
		Modified time.Time     `bson:"modified" json:"modified"`
	}
)

func (iAmAlive *IAmAlive) GetId() bson.ObjectId {
	return iAmAlive.Id
}

func (iAmAlive *IAmAlive) SetId(id bson.ObjectId) {
	iAmAlive.Id = id
}

func (iAmAlive *IAmAlive) SetCreated(time time.Time) {
	iAmAlive.Created = time
}

func (iAmAlive *IAmAlive) SetModified(time time.Time) {
	iAmAlive.Modified = time
}

func (iAmAlive *IAmAlive) SetContent(content string) {
	iAmAlive.Content = content
}

func (iAmAlive *IAmAlive) GetContent() string {
	return iAmAlive.Content
}

func (iAmAlive *IAmAlive) getEntity() IAmAlive {
	return *iAmAlive
}

