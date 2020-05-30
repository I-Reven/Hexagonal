package rest

import (
	"github.com/I-Reven/Hexagonal/src/applications/core/service"
	"github.com/I-Reven/Hexagonal/src/infrastructures/logger"
	"github.com/gin-gonic/gin"
	"github.com/juju/errors"
	"net/http"
)

type Tracker struct {
	log     logger.Log
	service service.TrackService
}

func (h *Tracker) Handler(context *gin.Context) {
	Track, err := h.service.GetTrack(context.Param("trackId"))

	if err != nil {
		err = errors.NewNotSupported(err, "error.handler-get-error-from-get-track-service")
		h.log.Error(err)
	}

	h.log.TraceLn(Track)
	context.JSON(http.StatusOK, Track)
}