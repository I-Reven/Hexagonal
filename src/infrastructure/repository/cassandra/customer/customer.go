package customer

import (
	"github.com/I-Reven/Hexagonal/src/framework/logger"
	"github.com/I-Reven/Hexagonal/src/infrastructure/repository/cassandra"
	"github.com/gocql/gocql"
	"github.com/juju/errors"
	"os"
)

var (
	session *gocql.Session
)

type Customer struct {
	log logger.Log
}

func (r *Customer) cql(keySpace string) *gocql.Session {
	cassandraConfig := cassandra.Cassandra{
		Host:        os.Getenv("ELASSANDRA_HOST"),
		Port:        os.Getenv("CASSANDRA_PORT"),
		KeySpace:    keySpace,
		Consistency: os.Getenv("CASSANDRA_CONSISTANCY_TRACKER"),
	}

	session = cassandraConfig.InitSession()

	return session
}

func (r *Customer) close() {
	session.Close()
}

func (r *Customer) MigrateKeySpace(keySpace string) error {
	cassandraConfig := cassandra.Cassandra{
		Host:        os.Getenv("ELASSANDRA_HOST"),
		Port:        os.Getenv("CASSANDRA_PORT"),
		KeySpace:    keySpace,
		Consistency: os.Getenv("CASSANDRA_CONSISTANCY_TRACKER"),
	}

	return cassandraConfig.MackKeySpace(keySpace)
}

func (r *Customer) MigrateMessage(keySpace string) error {
	messageQuery := `CREATE TYPE IF NOT EXISTS message (
						id TEXT,
						userId BIGINT,
  						content TEXT,
						kind INT,
  						seen SET<BIGINT>,
						delivered SET<BIGINT>,
						timestamp BIGINT,
					);`

	return r.cql(keySpace).Query(messageQuery).Exec()
}

func (r *Customer) MigrateMetaData(keySpace string) error {
	messageQuery := `CREATE TYPE IF NOT EXISTS metaData (
  						key TEXT,
						kind INT,
  						value TEXT,
					);`

	return r.cql(keySpace).Query(messageQuery).Exec()
}

func (r *Customer) MigrateRoom(keySpace string) error {
	roomQuery := `CREATE TABLE IF NOT EXISTS rooms (
  					id TIMEUUID,
					roomId BIGINT,
  					status INT,
  					usersId SET<BIGINT>,
					messages SET<frozen <message>>,
					metaData SET<frozen <metaData>>,
  					rating INT,
  					timestamp TIMESTAMP,
  					PRIMARY KEY(id)
				);`

	return r.cql(keySpace).Query(roomQuery).Exec()
}

func (r *Customer) Migrate(keySpace string) error {

	if err := r.MigrateKeySpace(keySpace); err != nil {
		err = errors.NewAlreadyExists(err, "error.can-not-migrate-this-key-space")
		r.log.Error(err)
		return err
	}

	if err := r.MigrateMessage(keySpace); err != nil {
		err = errors.NewAlreadyExists(err, "error.can-not-migrate-message")
		r.log.Error(err)
		return err
	}

	if err := r.MigrateMetaData(keySpace); err != nil {
		err = errors.NewAlreadyExists(err, "error.can-not-migrate-meta-data")
		r.log.Error(err)
		return err
	}

	if err := r.MigrateRoom(keySpace); err != nil {
		err = errors.NewAlreadyExists(err, "error.can-not-migrate-room")
		r.log.Error(err)
		return err
	}

	return nil
}
