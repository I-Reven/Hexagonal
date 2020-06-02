package entity

import (
	"fmt"
	"github.com/bxcodec/faker/v3"
	"github.com/gocql/gocql"
	"reflect"
	"time"
)

type Track struct {
	Id        gocql.UUID `cql:"id" json:"id" faker:"-"`
	TrackId   gocql.UUID `cql:"track_id" json:"trackId" faker:"-"`
	Message   string     `cql:"message" json:"message" faker:"sentence"`
	Data      []string   `cql:"data" json:"data"`
	Debugs    []Debug    `cql:"debugger" json:"debugger"`
	Error     string     `cql:"error" json:"error" faker:"error"`
	Timestamp time.Time  `cql:"timestamp" json:"timestamp"`
}

func (e *Track) GetId() gocql.UUID                { return e.Id }
func (e *Track) SetId(id gocql.UUID)              { e.Id = id }
func (e *Track) GetTrackId() gocql.UUID           { return e.TrackId }
func (e *Track) SetTrackId(trackId gocql.UUID)    { e.TrackId = trackId }
func (e *Track) GetMessage() string               { return e.Message }
func (e *Track) SetMessage(message string)        { e.Message = message }
func (e *Track) GetData() []string                { return e.Data }
func (e *Track) SetData(data []string)            { e.Data = data }
func (e *Track) AddData(data interface{})         { e.Data = append(e.Data, fmt.Sprintf("%v", data)) }
func (e *Track) GetError() string                 { return e.Error }
func (e *Track) SetError(err string)              { e.Error = err }
func (e *Track) GetTimestamp() time.Time          { return e.Timestamp }
func (e *Track) SetTimestamp(timestamp time.Time) { e.Timestamp = timestamp }
func (e *Track) GetDebugs() []Debug               { return e.Debugs }
func (e *Track) SetDebugs(debugs []Debug)         { e.Debugs = debugs }

func (e *Track) AddDebug(message string, data ...interface{}) {
	e.Debugs = append(e.Debugs, CreateDebugger(message, data...))
}

func (e *Track) Factory() error { e.customFaker(); return faker.FakeData(e) }
func (e *Track) customFaker() {
	_ = faker.AddProvider("error", func(v reflect.Value) (interface{}, error) {
		return "error.faker.error-message", nil
	})
}
