package middleware

import (
	"github.com/I-Reven/Hexagonal/src/applications/core/service"
	"github.com/I-Reven/Hexagonal/src/infrastructures/logger"
	"github.com/gin-gonic/gin"
	"os"
)

type Tracker struct {
	Log   logger.Log
	Track logger.Tracker
}

func (m Tracker) RequestTracker() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := m.Track.Create(ctx)
		if err != nil {
			ctx.Next()
		} else {
			if os.Getenv("DEBUG_KEY") == ctx.GetHeader("Debug") {
				_ = m.Log.StartDebug()
			}

			m.Log.TraceLn(ctx.Request)

			ctx.Set("TrackerId", id)
			ctx.Header("tracker-id", id)
			ctx.Next()

			service.TrackService{}.TrackRequestProducer(id)

			if os.Getenv("DEBUG_KEY") == ctx.GetHeader("Debug") {
				_ = m.Log.EndDebug()
			}
		}
	}
}
