package service

import (
	"github.com/I-Reven/Hexagonal/src/domains/entity"
	message "github.com/I-Reven/Hexagonal/src/domains/message/rabbit"
	"github.com/I-Reven/Hexagonal/src/infrastructures/logger"
	"github.com/I-Reven/Hexagonal/src/infrastructures/queue/rabbit"
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

func TrackRequestProducer(id string) {
	mes := message.TrackRequest{
		Id: id,
	}

	err := rabbit.ProduceMessage(mes)

	if err != nil {
		err = errors.NewNotSupported(err, "error.service-can-not-producer-track-request")
		logger.Warn(err)
	}

	logger.Debug("requestTrackerService.TrackRequestProducer", mes)
}
