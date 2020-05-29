package rest

import (
	"github.com/I-Reven/Hexagonal/src/applications/core/service"
	"github.com/I-Reven/Hexagonal/src/infrastructures/logger"
	"github.com/gin-gonic/gin"
	"github.com/juju/errors"
	"net/http"
)

func Ping(ctx *gin.Context) {
	ctx.String(http.StatusOK, "PONG")
}

func IAmAlive(ctx *gin.Context) {
	iAmAlive, err := service.GetLastTest()

	if err != nil {
		err = errors.NewNotSupported(err, "error.Handler-get-error-from-test-service")
		logger.Error(err)
	}

	service.Test()

	ctx.JSON(http.StatusOK, iAmAlive)
}
