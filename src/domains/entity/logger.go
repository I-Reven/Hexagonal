package entity

import (
	"github.com/gocql/gocql"
	"time"
)

type Logger struct {
	Id        gocql.UUID `cql:"id"`
	Message   string     `cql:"message"`
	Data      []string   `cql:"data"`
	Error     error      `cql:"error"`
	Timestamp time.Time  `cql:"timestamp"`
}

func (l Logger) GetId() gocql.UUID                { return l.Id }
func (l Logger) SetId(id gocql.UUID)              { l.Id = id }
func (l Logger) GetMessage() string               { return l.Message }
func (l Logger) SetMessage(message string)        { l.Message = message }
func (l Logger) GetData() []string                { return l.Data }
func (l Logger) SetData(data []string)            { l.Data = data }
func (l Logger) GetError() error                  { return l.Error }
func (l Logger) SetError(err error)               { l.Error = err }
func (l Logger) GetTimestamp() time.Time          { return l.Timestamp }
func (l Logger) SetTimestamp(timestamp time.Time) { l.Timestamp = timestamp }
