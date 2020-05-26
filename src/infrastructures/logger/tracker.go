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

func CreateTracker(context *gin.Context) (string, error) {
	Session = sessions.Default(context)
	id, err := track.CreateTrack()

	if err != nil {
		return "", err
	}
	Session.Set("track-id", id)

	return id, Session.Save()
}

func TrackMessage(message string) {
	if Session != nil {
		id := Session.Get("track-id").(string)
		err := track.AddMessage(id, message)

		if err != nil {
			Error(err)
		}
	}
}

func TrackError(error error) {
	if Session != nil {
		id := Session.Get("track-id").(string)
		err := track.AddError(id, error)

		if err != nil {
			Error(err)
		}
	}
}

func TrackData(data ...interface{}) {
	if Session != nil {
		id := Session.Get("track-id").(string)

		for _, info := range data {
			err := track.AddData(id, info)

			if err != nil {
				Error(err)
			}
		}
	}
}

func TrackDebug(message string, data ...interface{}) {
	if Session != nil {
		id := Session.Get("track-id").(string)

		err := track.AddDebug(id, message, data...)

		if err != nil {
			Error(err)
		}
	}
}

func GetTrack() entity.Track {
	if Session != nil {
		id := Session.Get("track-id").(string)
		t, err := track.GetTrack(id)

		if err != nil {
			Error(err)
		}

		return t
	}

	return entity.Track{}
}
