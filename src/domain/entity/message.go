package entity

import (
	"github.com/bxcodec/faker/v3"
	"github.com/gocql/gocql"
)

type Message struct {
	Id        gocql.UUID `cql:"id" json:"id" faker:"-"`
	UserId    int64      `cql:"userId" json:"user_id"`
	Content   string     `cql:"content" json:"content" faker:"sentence"`
	Kind      int32      `cql:"kind" json:"kind"`
	Seen      []int64    `cql:"seen" json:"seen"`
	Delivered []int64    `cql:"delivered" json:"delivered"`
	Timestamp int64      `cql:"timestamp" json:"timestamp"`
}

func (e *Message) GetId() gocql.UUID                       { return e.Id }
func (e *Message) SetId(id gocql.UUID) *Message            { e.Id = id; return e }
func (e *Message) MakeId() *Message                        { e.SetId(gocql.TimeUUID()); return e }
func (e *Message) GetUserId() int64                        { return e.UserId }
func (e *Message) SetUserId(userId int64) *Message         { e.UserId = userId; return e }
func (e *Message) GetContent() string                      { return e.Content }
func (e *Message) SetContent(content string) *Message      { e.Content = content; return e }
func (e *Message) GetKind() int32                          { return e.Kind }
func (e *Message) SetKind(kind int32) *Message             { e.Kind = kind; return e }
func (e *Message) GetSeen() []int64                        { return e.Seen }
func (e *Message) SetSeen(seen []int64) *Message           { e.Seen = seen; return e }
func (e *Message) GetDelivered() []int64                   { return e.Delivered }
func (e *Message) SetDelivered(delivered []int64) *Message { e.Delivered = delivered; return e }
func (e *Message) GetTimestamp() int64                     { return e.Timestamp }
func (e *Message) SetTimestamp(timestamp int64) *Message   { e.Timestamp = timestamp; return e }
func (e *Message) AddSeen(userId int64) *Message           { e.SetSeen(append(e.GetSeen(), userId)); return e }
func (e *Message) AddDelivered(userId int64) *Message {
	e.SetDelivered(append(e.GetDelivered(), userId))
	return e
}
func (e *Message) Factory() error { return faker.FakeData(e) }

func (e *Message) Create(content string, kind int32, userId int64) {
	e.MakeId()
	e.SetContent(content)
	e.SetKind(kind)
	e.SetUserId(userId)
	e.AddSeen(userId)
	e.AddDelivered(userId)
}

func (e *Message) Make(content string, kind int32, userId int64) {
	e.SetContent(content)
	e.SetKind(kind)
	e.SetUserId(userId)
	e.AddSeen(userId)
	e.AddDelivered(userId)
}