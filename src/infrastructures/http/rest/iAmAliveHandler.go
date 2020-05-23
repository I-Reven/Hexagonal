package rest

import (
	"github.com/I-Reven/Hexagonal/src/applications/core/service"
	"github.com/I-Reven/Hexagonal/src/infrastructures/logger"
	"github.com/gin-gonic/gin"
	"github.com/juju/errors"
	"net/http"
)

func IAmAlive(context *gin.Context) {
	err, iAmAlive := service.GetLastTest()

	if err != nil {
		err = errors.NewNotSupported(err, "error.Handler-get-error-from-test-service")
		logger.Error(err)
	} else {
		service.Test()
	}

	context.JSON(http.StatusOK, iAmAlive)
}
