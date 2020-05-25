package log

import (
	"github.com/I-Reven/Hexagonal/src/domains/entity"
	"github.com/I-Reven/Hexagonal/src/infrastructures/repository/cassandra"
	"github.com/gocql/gocql"
	"github.com/juju/errors"
	"os"
	"time"
)

type (
	Logger struct {
		*gocql.Session
	}
)

func Log() Logger {
	cassandraConfig := cassandra.Cassandra{
		Host:        os.Getenv("CASSANDRA_HOST"),
		Port:        os.Getenv("CASSANDRA_PORT"),
		Keyspace:    os.Getenv("CASSANDRA_KEYSPACE_LOGGER"),
		Consistancy: os.Getenv("CASSANDRA_CONSISTANCY_LOGGER"),
	}

	return Logger{cassandraConfig.InitSession()}
}

func (l Logger) Migrate() error {
	query := `
		CREATE TABLE IF NOT EXISTS logs (
  			id TIMEUUID,
  			message TEXT,
  			data SET<TEXT>,
  			error TEXT,
  			timestamp TIMESTAMP,
  			PRIMARY KEY(id)
		);
    	`
	return l.Query(query).Exec()
}

func (l Logger) Create(log *entity.Log) error {
	log.Id = gocql.TimeUUID()
	query := `
		INSERT INTO logs (
		    id,
		    message,
		    data,
		    error,
			timestamp
		)
		VALUES (?, ?, ?, ?, ?)
    	`
	return l.Query(query,
		log.GetId(),
		log.GetMessage(),
		log.GetData(),
		log.GetError(),
		time.Now()).Exec()
}

func (l Logger) GetById(id gocql.UUID) (*entity.Log, error) {
	m := map[string]interface{}{}
	query := `
		SELECT * FROM logger
			WHERE id = ?
		LIMIT 1
    	`
	itr := l.Query(query, id).Consistency(gocql.One).Iter()
	for itr.MapScan(m) {
		log := &entity.Log{}
		log.SetId(m["id"].(gocql.UUID))
		log.SetMessage(m["message"].(string))
		log.SetData(m["data"].([]string))
		log.SetError(m["error"].(string))
		log.SetTimestamp(m["timestamp"].(time.Time))

		return log, nil
	}

	return nil, errors.NewNotFound(errors.New("error"), "Log Not Found")
}

func (l Logger) Update(id gocql.UUID, message string, error string) error {
	query := `
        	UPDATE logs
		SET message = ?, error = ?
		WHERE id = ?
    	`
	return l.Query(query, message, error, id).Exec()
}

func (l Logger) AddData(id gocql.UUID, data string) error {
	query := `
		UPDATE logs
		SET data = data + ?
		WHERE id = ?;
	`
	return l.Query(query, []string{data}, id).Exec()
}

func (l Logger) RemoveData(company string, id gocql.UUID, data string) error {
	query := `
		UPDATE logs
		SET data = data - ?
		WHERE id = ?;
	`
	return l.Query(query, []string{data}, company, id).Exec()
}
