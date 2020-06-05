package cassandra

import (
	"github.com/I-Reven/Hexagonal/src/framework/logger"
	"github.com/I-Reven/Hexagonal/src/infrastructure/repository/cassandra/tracker"
	"github.com/juju/errors"
)

type Migration struct {
	log   logger.Log
	track tracker.Track
}

func (m *Migration) Migrate() error {
	if err := m.track.MigrateKeySpace(); err != nil {
		err = errors.NewNotSupported(err, "error.can-not-migrate-cassandra-tracks-key-space")
		return err
	}

	if err := m.track.Migrate(); err != nil {
		err = errors.NewNotSupported(err, "error.can-not-migrate-cassandra-tracks")
		m.log.Error(err)
		return err
	}

	return nil
}
