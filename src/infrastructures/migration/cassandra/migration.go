package cassandra

import (
	"github.com/I-Reven/Hexagonal/src/infrastructures/logger"
	"github.com/I-Reven/Hexagonal/src/infrastructures/repository/cassandra/track"
	"github.com/juju/errors"
)

type Migration struct {
	Log logger.Log
}

func (m Migration) Migrate() error {
	err := track.Track().Migrate()

	if err != nil {
		err = errors.NewNotSupported(err, "error.can-not-migrate-cassandra-tracks")
		m.Log.Error(err)
	}

	return err
}
