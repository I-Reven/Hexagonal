package entity

import (
	"github.com/gocql/gocql"
	"time"
)

type Log struct {
	Id        gocql.UUID `cql:"id"`
	Message   string     `cql:"message"`
	Data      []string   `cql:"data"`
	Error     string     `cql:"error"`
	Timestamp time.Time  `cql:"timestamp"`
}

func (l Log) GetId() gocql.UUID                { return l.Id }
func (l Log) SetId(id gocql.UUID)              { l.Id = id }
func (l Log) GetMessage() string               { return l.Message }
func (l Log) SetMessage(message string)        { l.Message = message }
func (l Log) GetData() []string                { return l.Data }
func (l Log) SetData(data []string)            { l.Data = data }
func (l Log) GetError() string                 { return l.Error }
func (l Log) SetError(err string)              { l.Error = err }
func (l Log) GetTimestamp() time.Time          { return l.Timestamp }
func (l Log) SetTimestamp(timestamp time.Time) { l.Timestamp = timestamp }
