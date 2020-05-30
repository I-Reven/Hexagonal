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
		Log     logger.Log
	}
)

func (j RequestTracker) Init(b []byte) (error, job.Job) {
	return json.Unmarshal(b, &j.Message), j
}

func (j RequestTracker) Handler() error {
	t, err := track.GetTrack(j.Message.Id)

	if err != nil {
		err = errors.NewNotSupported(err, "error.request-tracker-can-not-find-track-data")
		j.Log.TraceLn(j.Message.Id)
		j.Log.Error(err)
		return err
	}

	tracker := repository.Track()
	err = tracker.Create(&t)

	if err != nil {
		err = errors.NewNotSupported(err, "error.request-tracker-can-not-save")
		j.Log.TraceLn(t)
		j.Log.Error(err)
		return err
	}

	for _, debug := range t.GetDebugs() {
		err = tracker.AddDebug(t.GetId(), debug)

		if err != nil {
			err = errors.NewNotSupported(err, "error.request-tracker-can-not-save-debug")
			j.Log.TraceLn(t, debug)
			j.Log.Error(err)
		}
	}

	return nil
}

func (i RequestTracker) Failed(err error) {
	err = errors.NewNotSupported(err, "error.job-request-tracker-failed")
	i.Log.TraceLn(i)
	i.Log.Fatal(err)
}

func (i RequestTracker) GetConfig() job.Config {
	return job.Config{Tries: i.Tries}
}
