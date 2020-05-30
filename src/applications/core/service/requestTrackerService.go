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

type TrackService struct {
	Log     logger.Log
	Produce rabbit.Produce
}

func (s TrackService) GetTrack(id string) (*entity.Track, error) {
	tracker := track.Track()
	Id, err := gocql.ParseUUID(id)

	if err != nil {
		err = errors.NewNotSupported(err, "error.can-not-parse-uuid")
		s.Log.Error(err)
		return nil, err
	}

	return tracker.GetByTrackId(Id)
}

func (s TrackService) TrackRequestProducer(id string) {
	mes := message.TrackRequest{
		Id: id,
	}

	err := s.Produce.ProduceMessage(mes)

	if err != nil {
		err = errors.NewNotSupported(err, "error.service-can-not-producer-track-request")
		s.Log.Warn(err)
	}

	s.Log.Debug("requestTrackerService.TrackRequestProducer", mes)
}
