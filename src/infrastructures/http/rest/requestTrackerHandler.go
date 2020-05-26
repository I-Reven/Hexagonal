package rest

import (
	"github.com/I-Reven/Hexagonal/src/applications/core/service"
	"github.com/I-Reven/Hexagonal/src/infrastructures/logger"
	"github.com/gin-gonic/gin"
	"github.com/juju/errors"
	"net/http"
)

func RequestTracker(context *gin.Context) {
	Track, err := service.GetTrack(context.Param("trackId"))

	if err != nil {
		err = errors.NewNotSupported(err, "error.handler-get-error-from-get-track-service")
		logger.Error(err)
	}

	logger.TraceLn(Track)
	context.JSON(http.StatusOK, Track)
}
