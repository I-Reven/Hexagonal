package job

import (
	"encoding/json"
	"github.com/I-Reven/Hexagonal/src/domains/job"
	"github.com/I-Reven/Hexagonal/src/domains/message/rabbit"
	"github.com/I-Reven/Hexagonal/src/infrastructures/logger"
	repository "github.com/I-Reven/Hexagonal/src/infrastructures/repository/cassandra/track"
	"github.com/I-Reven/Hexagonal/src/infrastructures/repository/redis/track"
	"github.com/juju/errors"
)

type (
	RequestTracker struct {
		Tries   int
		Message rabbit.TrackRequest
	}
)

func (i RequestTracker) Init(b []byte) (error, job.Job) {
	return json.Unmarshal(b, &i.Message), i
}

func (i RequestTracker) Handler() error {
	t, err := track.GetTrack(i.Message.Id)

	if err != nil {
		err = errors.NewNotSupported(err, "error.request-tracker-can-not-find-track-data")
		logger.TraceLn(i.Message.Id)
		logger.Error(err)
		return err
	}

	tracker := repository.Track()
	err = tracker.Create(&t)

	if err != nil {
		err = errors.NewNotSupported(err, "error.request-tracker-can-not-save")
		logger.TraceLn(t)
		logger.Error(err)
		return err
	}

	for _, debug := range t.GetDebugs() {
		err = tracker.AddDebug(t.GetId(), debug)

		if err != nil {
			err = errors.NewNotSupported(err, "error.request-tracker-can-not-save-debug")
			logger.TraceLn(t, debug)
			logger.Error(err)
		}
	}

	return nil
}

func (i RequestTracker) Failed(err error) {
	err = errors.NewNotSupported(err, "error.job-request-tracker-failed")
	logger.TraceLn(i)
	logger.Fatal(err)
}

func (i RequestTracker) GetConfig() job.Config {
	return job.Config{Tries: i.Tries}
}
