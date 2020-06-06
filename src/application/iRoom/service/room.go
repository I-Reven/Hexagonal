package service

import (
	"github.com/I-Reven/Hexagonal/src/domain/entity"
	"github.com/I-Reven/Hexagonal/src/infrastructure/repository/cassandra/customer"
)

type Room struct {
	repository customer.Room
	entity entity.Room

}

func (s Room)Create(customer string, roomId int64, userId int64) error {
	return s.repository.Create(customer, s.entity.Make(roomId, userId))
}
