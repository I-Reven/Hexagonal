package service

import (
	"github.com/I-Reven/Hexagonal/src/domains/entity"
	"github.com/I-Reven/Hexagonal/src/infrastructures/logger"
	"github.com/I-Reven/Hexagonal/src/infrastructures/repository/cassandra/track"
	"github.com/gocql/gocql"
	"github.com/juju/errors"
)

func GetTrack(id string) (*entity.Track, error) {
	tracker := track.Track()
	Id, err := gocql.ParseUUID(id)

	if err != nil {
		err = errors.NewNotSupported(err, "error.can-not-parse-uuid")
		logger.Error(err)
		return nil, err
	}

	return tracker.GetByTrackId(Id)
}
