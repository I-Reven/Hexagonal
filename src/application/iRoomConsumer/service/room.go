package service

import (
	"github.com/I-Reven/Hexagonal/src/domain/entity"
	"github.com/I-Reven/Hexagonal/src/framework/logger"
	"github.com/I-Reven/Hexagonal/src/infrastructure/repository/cassandra/customer"
	"github.com/juju/errors"
)

type Room struct {
	repository customer.Room
	entity     *entity.Room
	message    *entity.Message
	log        logger.Log
}

func (s Room) Create(customer string, roomId int64, userId int64) error {

	if err := s.repository.Create(customer, s.entity.Make(roomId, userId)); err != nil {
		err = errors.NewNotSupported(err, "error.can-not-create-new-rom")
		s.log.Error(err)
		return err
	}

	return nil
}

func (s Room) AddUser(customer string, roomId int64, userId int64) error {
	var err error

	if s.entity, err = s.GetByRoomId(customer, roomId); err != nil {
		return err
	}

	if err = s.repository.AddUser(customer, s.entity.GetId(), userId); err != nil {
		err = errors.NewNotSupported(err, "error.can-not-add-user-to-room")
		s.log.Error(err)
		return err
	}

	return nil
}

func (s Room) AddMessage(customer string, roomId int64, userId int64, content string, kind int32) error {
	var err error

	if s.entity, err = s.GetByRoomId(customer, roomId); err != nil {
		return err
	}
	if err = s.repository.AddMessage(customer, s.entity.GetId(), *s.message.Make(content, kind, userId)); err != nil {
		err = errors.NewNotSupported(err, "error.can-not-add-message-to-room")
		s.log.Error(err)
		return err
	}

	return nil
}

func (s Room) GetByRoomId(customer string, roomId int64) (*entity.Room, error) {
	var err error

	if s.entity, err = s.repository.GetByRoomId(customer, roomId); err != nil {
		err = errors.NewNotFound(err, "error.can-not-found-room")
		s.log.Error(err)
		return nil, err
	}

	return s.entity, nil
}
