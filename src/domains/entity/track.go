package entity

import (
	"fmt"
	"github.com/gocql/gocql"
	"time"
)

type Track struct {
	Id        gocql.UUID `cql:"id" json:"id"`
	TrackId   gocql.UUID `cql:"track_id" json:"trackId"`
	Message   string     `cql:"message" json:"message"`
	Data      []string   `cql:"data" json:"data"`
	Debugs    []Debug    `cql:"debugger" json:"debugger"`
	Error     string     `cql:"error" json:"error"`
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
