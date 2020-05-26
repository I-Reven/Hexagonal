package track

import (
	"github.com/I-Reven/Hexagonal/src/domains/entity"
	"github.com/I-Reven/Hexagonal/src/infrastructures/repository/cassandra"
	"github.com/gocql/gocql"
	"github.com/juju/errors"
	"github.com/mitchellh/mapstructure"
	"os"
	"time"
)

type (
	Tracker struct {
		*gocql.Session
	}
)

var (
	toDebugs = func(i interface{}) []entity.Debug {
		var debugs []entity.Debug
		debug := entity.Debug{}

		for _, d := range i.([]map[string]interface{}) {
			err := mapstructure.Decode(d, &debug)
			if err == nil {
				debugs = append(debugs, debug)
			}
		}

		return debugs
	}
)

func Track() Tracker {
	cassandraConfig := cassandra.Cassandra{
		Host:        os.Getenv("CASSANDRA_HOST"),
		Port:        os.Getenv("CASSANDRA_PORT"),
		Keyspace:    os.Getenv("CASSANDRA_KEYSPACE_TRACKER"),
		Consistancy: os.Getenv("CASSANDRA_CONSISTANCY_TRACKER"),
	}

	return Tracker{cassandraConfig.InitSession()}
}

func (t Tracker) Migrate() error {
	debugQuery := `
		CREATE TYPE IF NOT EXISTS debug (
  			message TEXT,
  			data SET<TEXT>,
			memory TEXT,
			cpu TEXT,
			timestamp BIGINT,
		);
    	`
	tracksQuery := `
		CREATE TABLE IF NOT EXISTS tracks (
  			id TIMEUUID,
			track_id UUID,
  			message TEXT,
  			data SET<TEXT>,
			debugs SET<frozen <debug>>,
  			error TEXT,
  			timestamp TIMESTAMP,
  			PRIMARY KEY(id)
		);
    	`
	err := t.Query(debugQuery).Exec()

	if err == nil {
		err = t.Query(tracksQuery).Exec()
	}

	return err
}

func (t Tracker) Create(track *entity.Track) error {
	track.Id = gocql.TimeUUID()
	query := `
		INSERT INTO tracks (
		    id,
		    track_id,
		    message,
		    data,
		    error,
			timestamp
		)
		VALUES (?, ?, ?, ?, ?, ?)
    	`
	return t.Query(query,
		track.GetId(),
		track.GetTrackId(),
		track.GetMessage(),
		track.GetData(),
		track.GetError(),
		time.Now()).Exec()
}

func (t Tracker) GetById(id gocql.UUID) (*entity.Track, error) {
	m := map[string]interface{}{}
	query := `
		SELECT * FROM tracks
			WHERE id = ?
		LIMIT 1
    	`
	itr := t.Query(query, id).Consistency(gocql.One).Iter()
	for itr.MapScan(m) {
		track := &entity.Track{}
		track.SetId(m["id"].(gocql.UUID))
		track.SetTrackId(m["track_id"].(gocql.UUID))
		track.SetMessage(m["message"].(string))
		track.SetData(m["data"].([]string))
		track.SetDebugs(toDebugs(m["debugs"]))
		track.SetError(m["error"].(string))
		track.SetTimestamp(m["timestamp"].(time.Time))

		return track, nil
	}

	return nil, errors.NewNotFound(errors.New("error"), "error.track-not-found")
}

func (t Tracker) GetByTrackId(trackId gocql.UUID) (*entity.Track, error) {
	m := map[string]interface{}{}
	query := `
		SELECT * FROM tracks
			WHERE track_id = ?
		LIMIT 1
		ALLOW FILTERING
    	`
	itr := t.Query(query, trackId).Consistency(gocql.One).Iter()
	for itr.MapScan(m) {
		track := &entity.Track{}
		track.SetId(m["id"].(gocql.UUID))
		track.SetTrackId(m["track_id"].(gocql.UUID))
		track.SetMessage(m["message"].(string))
		track.SetData(m["data"].([]string))
		track.SetDebugs(toDebugs(m["debugs"]))
		track.SetError(m["error"].(string))
		track.SetTimestamp(m["timestamp"].(time.Time))

		return track, nil
	}

	return nil, errors.NewNotFound(errors.New("error"), "error.track-not-found")
}

func (t Tracker) Update(id gocql.UUID, message string, error string) error {
	query := `
        	UPDATE tracks
		SET message = ?, error = ?
		WHERE id = ?
    	`
	return t.Query(query, message, error, id).Exec()
}

func (t Tracker) AddData(id gocql.UUID, data string) error {
	query := `
		UPDATE tracks
		SET data = data + ?
		WHERE id = ?;
	`
	return t.Query(query, []string{data}, id).Exec()
}

func (t Tracker) RemoveData(id gocql.UUID, data string) error {
	query := `
		UPDATE tracks
		SET data = data - ?
		WHERE id = ?;
	`
	return t.Query(query, []string{data}, id).Exec()
}

func (t Tracker) AddDebug(id gocql.UUID, debug entity.Debug) error {
	query := `
		UPDATE tracks
		SET debugs = debugs + ?
		WHERE id = ?;
	`
	return t.Query(query, []entity.Debug{debug}, id).Exec()
}

func (t Tracker) RemoveDebug(id gocql.UUID, debug entity.Debug) error {
	query := `
		UPDATE tracks
		SET debugs = debugs - ?
		WHERE id = ?;
	`
	return t.Query(query, []entity.Debug{debug}, id).Exec()
}
