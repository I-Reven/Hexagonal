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

func (m *Message) GetId() gocql.UUID                       { return m.Id }
func (m *Message) SetId(id gocql.UUID) *Message            { m.Id = id; return m }
func (m *Message) MakeId() *Message                        { m.SetId(gocql.TimeUUID()); return m }
func (m *Message) GetUserId() int64                        { return m.UserId }
func (m *Message) SetUserId(userId int64) *Message         { m.UserId = userId; return m }
func (m *Message) GetContent() string                      { return m.Content }
func (m *Message) SetContent(content string) *Message      { m.Content = content; return m }
func (m *Message) GetKind() int32                          { return m.Kind }
func (m *Message) SetKind(kind int32) *Message             { m.Kind = kind; return m }
func (m *Message) GetSeen() []int64                        { return m.Seen }
func (m *Message) SetSeen(seen []int64) *Message           { m.Seen = seen; return m }
func (m *Message) GetDelivered() []int64                   { return m.Delivered }
func (m *Message) SetDelivered(delivered []int64) *Message { m.Delivered = delivered; return m }
func (m *Message) GetTimestamp() int64                     { return m.Timestamp }
func (m *Message) SetTimestamp(timestamp int64) *Message   { m.Timestamp = timestamp; return m }
func (m *Message) AddSeen(userId int64) *Message           { m.SetSeen(append(m.GetSeen(), userId)); return m }
func (m *Message) AddDelivered(userId int64) *Message {
	m.SetDelivered(append(m.GetDelivered(), userId))
	return m
}
func (m *Message) Factory() error { return faker.FakeData(m) }

func (m *Message) Create(content string, kind int32, userId int64) {
	m.MakeId()
	m.SetContent(content)
	m.SetKind(kind)
	m.SetUserId(userId)
	m.AddSeen(userId)
	m.AddDelivered(userId)
}

func (m *Message) Make(content string, kind int32, userId int64) {
	m.SetContent(content)
	m.SetKind(kind)
	m.SetUserId(userId)
	m.AddSeen(userId)
	m.AddDelivered(userId)
}
