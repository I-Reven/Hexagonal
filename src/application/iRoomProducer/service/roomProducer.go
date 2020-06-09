package service

import (
	"github.com/I-Reven/Hexagonal/src/domain/entity"
	message "github.com/I-Reven/Hexagonal/src/domain/message/rabbit"
	"github.com/I-Reven/Hexagonal/src/framework/logger"
	queue "github.com/I-Reven/Hexagonal/src/framework/queue/rabbit"
	"github.com/juju/errors"
)

type RoomProducer struct {
	log     logger.Log
	produce queue.Produce
	message entity.Message
}

func (s RoomProducer) Create(customerName string, roomId int64, userId int64) error {
	msg := message.CreateRoom{CustomerName: customerName, RoomId: roomId, UserId: userId}

	if err := s.produce.ProduceMessage(msg); err != nil {
		err = errors.NewNotSupported(err, "error.can-not-produce-create-room-message")
		s.log.Error(err)
		return err
	}

	return nil
}

func (s RoomProducer) AddUser(customerName string, roomId int64, userId int64) error {
	msg := message.AddUser{CustomerName: customerName, RoomId: roomId, UserId: userId}

	if err := s.produce.ProduceMessage(msg); err != nil {
		err = errors.NewNotSupported(err, "error.can-not-produce-add-user-message")
		s.log.Error(err)
		return err
	}

	return nil
}

func (s RoomProducer) AddMessage(customerName string, roomId int64, userId int64, content string, kind int32) (string, error) {
	msg := message.AddMessage{
		CustomerName: customerName,
		RoomId:       roomId,
		UserId:       userId,
		MessageId:    s.message.MakeId().GetId(),
		Content:      content,
		Kind:         kind,
	}

	if err := s.produce.ProduceMessage(msg); err != nil {
		err = errors.NewNotSupported(err, "error.can-not-produce-add-message-message")
		s.log.Error(err)
		return s.message.GetId(), err
	}

	return s.message.GetId(), nil
}

func (s RoomProducer) SeenMessage(customerName string, roomId int64, messageId string, userId int64) error {
	msg := message.SeenMessage{CustomerName: customerName, MessageId: messageId, RoomId: roomId, UserId: userId}

	if err := s.produce.ProduceMessage(msg); err != nil {
		err = errors.NewNotSupported(err, "error.can-not-produce-seen-message-message")
		s.log.Error(err)
		return err
	}

	return nil
}

func (s RoomProducer) DeliverMessage(customerName string, roomId int64, messageId string, userId int64) error {
	msg := message.DeliverMessage{CustomerName: customerName, MessageId: messageId, RoomId: roomId, UserId: userId}

	if err := s.produce.ProduceMessage(msg); err != nil {
		err = errors.NewNotSupported(err, "error.can-not-produce-seen-message-message")
		s.log.Error(err)
		return err
	}

	return nil
}

func (s RoomProducer) AddMetaData(customerName string, roomId int64, key string, kind int32, value string) error {
	msg := message.AddMetaData{CustomerName: customerName, RoomId: roomId, Key: key, Kind: kind, Value: value}

	if err := s.produce.ProduceMessage(msg); err != nil {
		err = errors.NewNotSupported(err, "error.can-not-produce-add-meta-data-message")
		s.log.Error(err)
		return err
	}

	return nil
}
