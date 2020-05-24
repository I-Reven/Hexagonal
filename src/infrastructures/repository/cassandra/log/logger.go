package log

import (
	"github.com/I-Reven/Hexagonal/src/domains/entity"
	"github.com/I-Reven/Hexagonal/src/infrastructures/repository/cassandra"
	"github.com/gocql/gocql"
	"github.com/juju/errors"
	"os"
	"sync"
	"time"
)

type (
	Logger struct {
		*gocql.Session
	}
)

var (
	cassandraConfig = cassandra.Cassandra{
		Host:        os.Getenv("CASSANDRA_HOST"),
		Port:        os.Getenv("CASSANDRA_PORT"),
		Keyspace:    os.Getenv("CASSANDRA_KEYSPACE_LOGGER"),
		Consistancy: os.Getenv("CASSANDRA_CONSISTANCY_LOGGER"),
	}

	logger Logger
	once   sync.Once
)

func Log() Logger {

	once.Do(func() {
		logger = Logger{cassandraConfig.InitSession()}
	})

	return logger
}

func (l Logger) Create(log *entity.Logger) error {
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
		log.Id,
		log.Message,
		log.Data,
		log.Error,
		log.Timestamp).Exec()
}

func (l Logger) GetById(id gocql.UUID) (*entity.Logger, error) {
	m := map[string]interface{}{}
	query := `
		SELECT * FROM logger
			WHERE id = ?
		LIMIT 1
    	`
	itr := l.Query(query, id).Consistency(gocql.One).Iter()
	for itr.MapScan(m) {
		log := &entity.Logger{}
		log.SetId(m["id"].(gocql.UUID))
		log.SetMessage(m["message"].(string))
		log.SetData(m["data"].([]string))
		log.SetError(m["error"].(error))
		log.SetTimestamp(m["timestamp"].(time.Time))

		return log, nil
	}

	return nil, errors.NewNotFound(errors.New("error"), "Log Not Found")
}

func (l Logger) Update(id gocql.UUID, message string, error error) error {
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
		SET tags = tags + ?
		WHERE id = ?;
	`
	return l.Query(query, []string{data}, id).Exec()
}

func (l Logger) RemoveData(company string, id gocql.UUID, tag string) error {
	query := `
		UPDATE logs
		SET tags = tags - ?
		WHERE id = ?;
	`
	return l.Query(query, []string{tag}, company, id).Exec()
}
