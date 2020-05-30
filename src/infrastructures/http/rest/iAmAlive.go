package rest

import (
	"github.com/I-Reven/Hexagonal/src/applications/core/service"
	"github.com/I-Reven/Hexagonal/src/infrastructures/logger"
	"github.com/gin-gonic/gin"
	"github.com/juju/errors"
	"net/http"
)

type IAmAlive struct {
	log     logger.Log
	service service.IAamAliveService
}

func (h *IAmAlive) Handler(ctx *gin.Context) {
	iAmAlive, err := h.service.GetLastTest()

	if err != nil {
		err = errors.NewNotSupported(err, "error.Handler-get-error-from-test-service")
		h.log.Error(err)
	}

	h.service.Test()

	ctx.JSON(http.StatusOK, iAmAlive)
}
