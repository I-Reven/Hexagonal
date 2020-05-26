package cassandra

import (
	"github.com/I-Reven/Hexagonal/src/infrastructures/logger"
	"github.com/I-Reven/Hexagonal/src/infrastructures/repository/cassandra/track"
	"github.com/juju/errors"
)

func Migrate() error {
	err := track.Track().Migrate()

	if err != nil {
		err = errors.NewNotSupported(err, "error.can-not-migrate-cassandra-tracks")
		logger.Error(err)
	}

	return err
}
