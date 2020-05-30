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
	log := logger.Log{}
	iAmAlive, err := service.IAamAliveService{}.GetLastTest()

	if err != nil {
		err = errors.NewNotSupported(err, "error.Handler-get-error-from-test-service")
		log.Error(err)
	}

	service.IAamAliveService{}.Test()

	ctx.JSON(http.StatusOK, iAmAlive)
}
