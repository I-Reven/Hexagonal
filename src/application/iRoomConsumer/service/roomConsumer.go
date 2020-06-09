package service

import (
	"github.com/I-Reven/Hexagonal/src/domain/entity"
	"github.com/I-Reven/Hexagonal/src/framework/logger"
	"github.com/I-Reven/Hexagonal/src/infrastructure/repository/cassandra/customer"
	"github.com/juju/errors"
)

type RoomConsumer struct {
	repository customer.Room
	log        logger.Log
}

func (s RoomConsumer) Create(customer string, roomId int64, userId int64) error {
	var err error
	room := &entity.Room{}

	if room, err = s.repository.GetByRoomId(customer, roomId); err != nil {
		room = &entity.Room{}
		if err = s.repository.Create(customer, room.Make(roomId, userId)); err != nil {
			err = errors.NewNotSupported(err, "error.can-not-create-new-room")
			s.log.Error(err)
			return err
		}
	} else {
		if err = s.repository.AddUser(customer, room.GetId(), userId); err != nil {
			err = errors.NewNotSupported(err, "error.can-not-update-room")
			s.log.Error(err)
			return err
		}
	}

	return nil
}

func (s RoomConsumer) AddUser(customer string, roomId int64, userId int64) error {
	var err error
	room := &entity.Room{}

	if room, err = s.GetByRoomId(customer, roomId); err != nil {
		return err
	}

	if err = s.repository.AddUser(customer, room.GetId(), userId); err != nil {
		err = errors.NewNotSupported(err, "error.can-not-add-user-to-room")
		s.log.Error(err)
		return err
	}

	return nil
}

func (s RoomConsumer) AddMessage(customer string, roomId int64, userId int64, messageId string, content string, kind int32) error {
	var (
		err     error
		message entity.Message
	)

	room := &entity.Room{}
	message.Make(messageId, content, kind, userId)

	if room, err = s.GetByRoomId(customer, roomId); err != nil {
		return err
	}

	if messageOld, exist := room.ExistMessage(message); !exist {
		if err = s.repository.AddMessage(customer, room.GetId(), message); err != nil {
			err = errors.NewNotSupported(err, "error.can-not-add-message-to-room")
			s.log.Error(err)
			return err
		}
	} else if err = s.repository.UpdateMessage(customer, room.GetId(), messageOld, message); err != nil {
		err = errors.NewNotSupported(err, "error.can-not-add-message-to-room")
		s.log.Error(err)
		return err
	}

	return nil
}

func (s RoomConsumer) SeenMessage(customer string, roomId int64, messageId string, userId int64) error {
	var (
		err     error
		message entity.Message
	)

	room := &entity.Room{}
	message.Make(messageId, "", 1, userId)

	if room, err = s.GetByRoomId(customer, roomId); err != nil {
		return err
	}

	if messageOld, exist := room.ExistMessage(message); exist {
		message = *messageOld.AddSeen(userId)

		if err = s.repository.UpdateMessage(customer, room.GetId(), messageOld, message); err != nil {
			err = errors.NewNotSupported(err, "error.can-not-update-message-seen")
			s.log.Error(err)
			return err
		}

		if _, exist := room.ExistUserId(userId); !exist {
			if err = s.AddUser(customer, roomId, userId); err != nil {
				return err
			}
		}

	} else {
		err = errors.NewNotSupported(err, "error.can-not-found-message")
	}

	return nil
}

func (s RoomConsumer) DeliverMessage(customer string, roomId int64, messageId string, userId int64) error {
	var (
		err     error
		message entity.Message
	)

	room := &entity.Room{}
	message.Make(messageId, "", 1, userId)

	if room, err = s.GetByRoomId(customer, roomId); err != nil {
		return err
	}

	if messageOld, exist := room.ExistMessage(message); exist {
		message = *messageOld.AddDelivered(userId)

		if err = s.repository.UpdateMessage(customer, room.GetId(), messageOld, message); err != nil {
			err = errors.NewNotSupported(err, "error.can-not-update-message-deliver")
			s.log.Error(err)
			return err
		}

		if _, exist := room.ExistUserId(userId); !exist {
			if err = s.AddUser(customer, roomId, userId); err != nil {
				return err
			}
		}

	} else {
		err = errors.NewNotSupported(err, "error.can-not-found-message")
	}

	return nil
}

func (s RoomConsumer) AddMetaData(customer string, roomId int64, key string, kind int32, value string) error {
	var (
		err      error
		metaData entity.MetaData
	)

	room := &entity.Room{}
	metaData.Make(key, value, kind)

	if room, err = s.GetByRoomId(customer, roomId); err != nil {
		return err
	}

	if metaDataOld, exist := room.ExistMetaData(metaData); !exist {
		if err = s.repository.AddMetaData(customer, room.GetId(), metaData); err != nil {
			err = errors.NewNotSupported(err, "error.can-not-add-message-to-rome")
			s.log.Error(err)
			return err
		}
	} else if err = s.repository.UpdateMetaData(customer, room.GetId(), metaDataOld, metaData); err != nil {
		err = errors.NewNotSupported(err, "error.can-not-update-message-to-rome")
		s.log.Error(err)
		return err
	}

	return nil
}

func (s RoomConsumer) GetByRoomId(customer string, roomId int64) (*entity.Room, error) {
	var err error
	room := &entity.Room{}

	if room, err = s.repository.GetByRoomId(customer, roomId); err != nil {
		err = errors.NewNotFound(err, "error.can-not-found-room")
		s.log.Error(err)
		return nil, err
	}

	return room, nil
}
