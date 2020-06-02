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
	err := m.track.Migrate()

	if err != nil {
		err = errors.NewNotSupported(err, "error.can-not-migrate-cassandra-tracks")
		m.log.Error(err)
	}

	return err
}
