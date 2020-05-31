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

func (t *Track) GetId() gocql.UUID                { return t.Id }
func (t *Track) SetId(id gocql.UUID)              { t.Id = id }
func (t *Track) GetTrackId() gocql.UUID           { return t.TrackId }
func (t *Track) SetTrackId(trackId gocql.UUID)    { t.TrackId = trackId }
func (t *Track) GetMessage() string               { return t.Message }
func (t *Track) SetMessage(message string)        { t.Message = message }
func (t *Track) GetData() []string                { return t.Data }
func (t *Track) SetData(data []string)            { t.Data = data }
func (t *Track) AddData(data interface{})         { t.Data = append(t.Data, fmt.Sprintf("%v", data)) }
func (t *Track) GetError() string                 { return t.Error }
func (t *Track) SetError(err string)              { t.Error = err }
func (t *Track) GetTimestamp() time.Time          { return t.Timestamp }
func (t *Track) SetTimestamp(timestamp time.Time) { t.Timestamp = timestamp }
func (t *Track) GetDebugs() []Debug               { return t.Debugs }
func (t *Track) SetDebugs(debugs []Debug)         { t.Debugs = debugs }

func (t *Track) AddDebug(message string, data ...interface{}) {
	t.Debugs = append(t.Debugs, CreateDebugger(message, data...))
}

func (t *Track) Factory() error { t.customFaker(); return faker.FakeData(t) }
func (t *Track) customFaker() {
	_ = faker.AddProvider("error", func(v reflect.Value) (interface{}, error) {
		return "error.faker.error-message", nil
	})
}
