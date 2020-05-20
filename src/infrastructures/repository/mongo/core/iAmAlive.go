package core

import (
	entity "github.com/I-Reven/Hexagonal/src/domains/entity/core"
	"gopkg.in/mgo.v2/bson"
)

type IAmAlive struct {
	entity.IAmAlive `bson:",inline"`
}

func (iAmAlive *IAmAlive) Add() error {
	iAmAlive.SetContent("I Am Alive")
	return iAmAlive.Save()
}

func (iAmAlive *IAmAlive) Save() error {
	return Mongo().Collection("iAmAlive").Save(iAmAlive)
}

func (iAmAlive *IAmAlive) GetById(Id string) error {
	return Mongo().Collection("iAmAlive").FindById(bson.ObjectId(Id), iAmAlive)
}
