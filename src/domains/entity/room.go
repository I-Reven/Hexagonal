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

func (r *Room) GetId() gocql.UUID                     { return r.Id }
func (r *Room) SetId(id gocql.UUID) *Room             { r.Id = id; return r }
func (r *Room) GetRoomId() int64                      { return r.RoomId }
func (r *Room) SetRoomId(roomId int64) *Room          { r.RoomId = roomId; return r }
func (r *Room) GetStatus() int32                      { return r.Status }
func (r *Room) SetStatus(status int32) *Room          { r.Status = status; return r }
func (r *Room) GetUsersId() []int64                   { return r.UsersId }
func (r *Room) SetUsersId(usersId []int64) *Room      { r.UsersId = usersId; return r }
func (r *Room) GetMessages() []Message                { return r.Messages }
func (r *Room) SetMessages(messages []Message) *Room  { r.Messages = messages; return r }
func (r *Room) GetRating() int32                      { return r.Rating }
func (r *Room) SetRating(rating int32) *Room          { r.Rating = rating; return r }
func (r *Room) Factory() error                        { return faker.FakeData(r) }
func (r *Room) GetMetaData() []MetaData               { return r.MetaData }
func (r *Room) SetMetaData(metaData []MetaData) *Room { r.MetaData = metaData; return r }

func (r *Room) AddMessage(message Message) *Room {
	r.AddUserId(message.GetUserId())
	r.SetMessages(append(r.GetMessages(), message))
	return r
}

func (r *Room) AddUserId(userId int64) *Room {
	var id int64
	for id = range r.GetUsersId() {
		if userId == id {
			return nil
		}
	}

	r.SetUserId(append(r.GetUsersId(), userId))
	return r
}

func (r *Room) AddMetaData(metaData MetaData) *Room {
	var (
		m     MetaData
		md    []MetaData
		exist = false
	)

	for m = range r.GetMetaData() {
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

	r.SetMetaData(md)
	return r
}

func (r *Room) Make(roomId int64, userId int64) *Room {
	r.SetRoomId(roomId)
	r.AddUserId(userId)
	return r
}
