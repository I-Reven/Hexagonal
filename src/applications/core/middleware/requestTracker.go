package middleware

import (
	"github.com/I-Reven/Hexagonal/src/infrastructures/logger"
	"github.com/I-Reven/Hexagonal/src/infrastructures/repository/cassandra/track"
	"github.com/gin-gonic/gin"
	"github.com/juju/errors"
	"os"
)

func RequestTracker() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := logger.CreateTracker(c)
		if err != nil {
			c.Next()
		} else {
			if os.Getenv("DEBUG_KEY") == c.GetHeader("Debug") {
				_ = logger.StartDebug()
			}

			logger.TraceLn(c.Request)

			c.Set("TrackerId", id)
			c.Header("tracker-id", id)
			c.Next()

			t := logger.GetTrack()
			tracker := track.Track()
			err = tracker.Create(&t)

			if err != nil {
				err = errors.NewNotSupported(err, "error.request-tracker-can-not-save")
				logger.Error(err)
			}

			for _, debug := range t.GetDebugs() {
				err = tracker.AddDebug(t.GetId(), debug)

				if err != nil {
					err = errors.NewNotSupported(err, "error.request-tracker-can-not-save-debug")
					logger.Error(err)
				}
			}

			if os.Getenv("DEBUG_KEY") == c.GetHeader("Debug") {
				_ = logger.EndDebug()
			}
		}
	}
}
