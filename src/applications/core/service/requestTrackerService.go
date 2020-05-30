package service

import (
	"github.com/I-Reven/Hexagonal/src/domains/entity"
	message "github.com/I-Reven/Hexagonal/src/domains/message/rabbit"
	"github.com/I-Reven/Hexagonal/src/infrastructures/logger"
	"github.com/I-Reven/Hexagonal/src/infrastructures/queue/rabbit"
	tracker2 "github.com/I-Reven/Hexagonal/src/infrastructures/repository/cassandra/tracker"
	"github.com/gocql/gocql"
	"github.com/juju/errors"
)

type TrackService struct {
	log     logger.Log
	produce rabbit.Produce
	track   tracker2.Track
}

func (s TrackService) GetTrack(id string) (*entity.Track, error) {
	Id, err := gocql.ParseUUID(id)

	if err != nil {
		err = errors.NewNotSupported(err, "error.can-not-parse-uuid")
		s.log.Error(err)
		return nil, err
	}

	return s.track.GetByTrackId(Id)
}

func (s TrackService) TrackRequestProducer(id string) {
	mes := message.TrackRequest{
		Id: id,
	}

	err := s.produce.ProduceMessage(mes)

	if err != nil {
		err = errors.NewNotSupported(err, "error.service-can-not-producer-track-request")
		s.log.Warn(err)
	}

	s.log.Debug("requestTrackerService.TrackRequestProducer", mes)
}
