package service

import (
	message "github.com/I-Reven/Hexagonal/domains/message/rabbit/core"
	"github.com/I-Reven/Hexagonal/infrastructures/queue/rabbit"
	repository "github.com/I-Reven/Hexagonal/infrastructures/repository/mongo/core"
)

func Test() error {
	err, iAmAlive := TestDatabase()

	if err != nil {
		return err
	}

	err = TestProducer(iAmAlive)

	return err
}

func TestDatabase() (error, repository.IAmAlive) {
	iAmAlive := repository.IAmAlive{}
	return iAmAlive.Add(), iAmAlive
}

func TestProducer(iAmAlive repository.IAmAlive) error {
	mes := message.IAmAlive{
		Id:      string(iAmAlive.GetId()),
		Content: iAmAlive.GetContent(),
	}

	return rabbit.ProduceMessage(mes)
}
