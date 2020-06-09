package entity

import (
	"github.com/bxcodec/faker/v3"
	"github.com/gocql/gocql"
)

type Room struct {
	Id       gocql.UUID `cql:"id" json:"id" faker:"-"`
	RoomId   int64      `cql:"room_id" json:"room_id"`
	Status   int32      `cql:"status" json:"status"`
	UsersId  []int64    `cql:"users_id" json:"users_id"`
	Messages []Message  `cql:"messages" json:"messages"`
	MetaData []MetaData `cql:"meta_data" json:"meta_data"`
	Rating   int32      `cql:"rating" json:"rating"`
}

func (e *Room) GetId() gocql.UUID                     { return e.Id }
func (e *Room) SetId(id gocql.UUID) *Room             { e.Id = id; return e }
func (e *Room) GetRoomId() int64                      { return e.RoomId }
func (e *Room) SetRoomId(roomId int64) *Room          { e.RoomId = roomId; return e }
func (e *Room) GetStatus() int32                      { return e.Status }
func (e *Room) SetStatus(status int32) *Room          { e.Status = status; return e }
func (e *Room) GetUsersId() []int64                   { return e.UsersId }
func (e *Room) SetUsersId(usersId []int64) *Room      { e.UsersId = usersId; return e }
func (e *Room) GetMessages() []Message                { return e.Messages }
func (e *Room) SetMessages(messages []Message) *Room  { e.Messages = messages; return e }
func (e *Room) GetRating() int32                      { return e.Rating }
func (e *Room) SetRating(rating int32) *Room          { e.Rating = rating; return e }
func (e *Room) Factory() error                        { return faker.FakeData(e) }
func (e *Room) GetMetaData() []MetaData               { return e.MetaData }
func (e *Room) SetMetaData(metaData []MetaData) *Room { e.MetaData = metaData; return e }

func (e *Room) AddMessage(message Message) *Room {
	e.AddUserId(message.GetUserId())
	e.SetMessages(append(e.GetMessages(), message))
	return e
}

func (e *Room) AddUserId(userId int64) *Room {
	e.SetUsersId(append(e.GetUsersId(), userId))
	return e
}

func (e *Room) AddMetaData(metaData MetaData) *Room {
	e.SetMetaData(append(e.GetMetaData(), metaData))
	return e
}

func (e *Room) ExistUserId(userId int64) (int64, bool) {
	for _, user := range e.GetUsersId() {
		if userId == user {
			return user, true
		}
	}

	return userId, false
}

func (e *Room) ExistMetaData(metaData MetaData) (MetaData, bool) {
	for _, md := range e.GetMetaData() {
		if md.Key == metaData.Key {
			return md, true
		}
	}

	return metaData, false
}

func (e *Room) ExistMessage(message Message) (Message, bool) {
	for _, msg := range e.GetMessages() {
		if msg.Id == message.Id {
			return msg, true
		}
	}

	return message, false
}

func (e *Room) Make(roomId int64, userId int64) *Room {
	e.SetRoomId(roomId)
	e.AddUserId(userId)
	e.SetStatus(1)
	return e
}
