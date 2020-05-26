package middleware

import (
	"github.com/I-Reven/Hexagonal/src/applications/core/service"
	"github.com/I-Reven/Hexagonal/src/infrastructures/logger"
	"github.com/gin-gonic/gin"
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

			service.TrackRequestProducer(id)

			if os.Getenv("DEBUG_KEY") == c.GetHeader("Debug") {
				_ = logger.EndDebug()
			}
		}
	}
}
