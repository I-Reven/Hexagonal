package core

import (
	"github.com/I-Reven/Hexagonal/src/applications/core/service"
	request "github.com/I-Reven/Hexagonal/src/domain/http"
	"github.com/I-Reven/Hexagonal/src/framework/logger"
	"github.com/gin-gonic/gin"
	"github.com/juju/errors"
	"net/http"
)

type Tracker struct {
	log     logger.Log
	service service.TrackService
	track   request.Track
}

func (h *Tracker) Handler(ctx *gin.Context) {
	if err := ctx.ShouldBindJSON(&h.track); err != nil {
		err = errors.NewBadRequest(err, "error.can-not-un-marshal-json")
		h.log.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.track.Validate(); err != nil {
		err = errors.NewNotValid(err, "error.request is not valid")
		h.log.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	Track, err := h.service.Get(h.track.TrackId)

	if err != nil {
		err = errors.NewNotSupported(err, "error.handler-get-error-from-get-track-service")
		h.log.Error(err)
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, Track)
}
