package entity

import (
	"github.com/bxcodec/faker/v3"
	"github.com/gocql/gocql"
)

type Room struct {
	Id       gocql.UUID `cql:"id" json:"id" faker:"-"`
	RoomId   int64      `cql:"roomId" json:"room_id"`
	Status   int32      `cql:"status" json:"status"`
	UsersId  []int64    `cql:"userId" json:"users_id"`
	Messages []Message  `cql:"messages" json:"messages"`
	MetaData []MetaData `cql:"metaData" json:"meta_data"`
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
	var id int64
	for id = range e.GetUsersId() {
		if userId == id {
			return nil
		}
	}

	e.SetUserId(append(e.GetUsersId(), userId))
	return e
}

func (e *Room) AddMetaData(metaData MetaData) *Room {
	var (
		m     MetaData
		md    []MetaData
		exist = false
	)

	for m = range e.GetMetaData() {
		if m.Key == metaData.Key {
			md = append(md, metaData)
			exist = true
		} else {
			md = append(md, m)
		}
	}

	if !exist {
		md = append(md, metaData)
	}

	e.SetMetaData(md)
	return e
}

func (e *Room) Make(roomId int64, userId int64) *Room {
	e.SetRoomId(roomId)
	e.AddUserId(userId)
	return e
}
