package logger

import (
	"github.com/I-Reven/Hexagonal/src/domains/entity"
	"github.com/I-Reven/Hexagonal/src/infrastructures/repository/redis/track"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

var (
	Session sessions.Session = nil
)

type Tracker struct {
	track track.Track
}

func (t *Tracker) Create(context *gin.Context) (string, error) {
	Session = sessions.Default(context)
	id, err := t.track.CreateTrack()

	if err != nil {
		return "", err
	}
	Session.Set("track-id", id)

	return id, Session.Save()
}

func (t *Tracker) Message(message string) error {
	if Session != nil {
		id := Session.Get("track-id").(string)
		return t.track.AddMessage(id, message)
	}
	return nil
}

func (t *Tracker) Error(error error) error {
	if Session != nil {
		id := Session.Get("track-id").(string)

		return t.track.AddError(id, error)
	}

	return nil
}

func (t *Tracker) Data(data ...interface{}) error {
	if Session != nil {
		id := Session.Get("track-id").(string)

		for _, info := range data {
			return t.track.AddData(id, info)
		}
	}

	return nil
}

func (t *Tracker) Debug(message string, data ...interface{}) error {
	if Session != nil {
		id := Session.Get("track-id").(string)

		return t.track.AddDebug(id, message, data...)
	}

	return nil
}

func (t *Tracker) Get() (entity.Track, error) {
	if Session != nil {
		id := Session.Get("track-id").(string)
		return t.track.GetTrack(id)
	}

	return entity.Track{}, nil
}
