package job

import (
	"encoding/json"
	"github.com/I-Reven/Hexagonal/src/domains/job"
	"github.com/I-Reven/Hexagonal/src/domains/message/rabbit"
	"github.com/I-Reven/Hexagonal/src/infrastructures/logger"
	"github.com/I-Reven/Hexagonal/src/infrastructures/repository/cassandra/tracker"
	redis "github.com/I-Reven/Hexagonal/src/infrastructures/repository/redis/track"
	"github.com/juju/errors"
)

type (
	RequestTracker struct {
		tries   int
		message rabbit.TrackRequest
		log     logger.Log
		track   tracker.Track
		redis   redis.Track
	}
)

func (j RequestTracker) Init(b []byte) (error, job.Job) {
	return json.Unmarshal(b, &j.message), j
}

func (j RequestTracker) Handler() error {
	t, err := j.redis.GetTrack(j.message.Id)

	if err != nil {
		err = errors.NewNotSupported(err, "error.request-tracker-can-not-find-track-data")
		j.log.TraceLn(j.message.Id)
		j.log.Error(err)
		return err
	}

	err = j.track.Create(&t)

	if err != nil {
		err = errors.NewNotSupported(err, "error.request-tracker-can-not-save")
		j.log.TraceLn(t)
		j.log.Error(err)
		return err
	}

	for _, debug := range t.GetDebugs() {
		err = j.track.AddDebug(t.GetId(), debug)

		if err != nil {
			err = errors.NewNotSupported(err, "error.request-tracker-can-not-save-debug")
			j.log.TraceLn(t, debug)
			j.log.Error(err)
		}
	}

	return nil
}

func (j RequestTracker) Failed(err error) {
	err = errors.NewNotSupported(err, "error.job-request-tracker-failed")
	j.log.TraceLn(j)
	j.log.Fatal(err)
}

func (j RequestTracker) GetConfig() job.Config {
	return job.Config{Tries: j.tries}
}
