package middleware

import (
	"github.com/I-Reven/Hexagonal/src/application/core/service"
	"github.com/I-Reven/Hexagonal/src/framework/logger"
	"github.com/gin-gonic/gin"
	"os"
)

type Tracker struct {
	log     logger.Log
	track   logger.Tracker
	service service.TrackService
}

func (m Tracker) Handler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := m.track.Create(ctx)
		if err != nil {
			ctx.Next()
		} else {
			if os.Getenv("DEBUG_KEY") == ctx.GetHeader("Debug") {
				_ = m.log.StartDebug()
			}

			m.log.TraceLn(ctx.Request)

			ctx.Set("TrackerId", id)
			ctx.Header("tracker-id", id)
			ctx.Next()

			m.service.Produce(id)

			if os.Getenv("DEBUG_KEY") == ctx.GetHeader("Debug") {
				_ = m.log.EndDebug()
			}
		}
	}
}
