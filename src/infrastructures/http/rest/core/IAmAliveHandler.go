package core

import (
	"github.com/I-Reven/Hexagonal/src/applications/core/service"
	"github.com/I-Reven/Hexagonal/src/infrastructures/logger"
	"github.com/juju/errors"
	"github.com/labstack/echo"
	"net/http"
)

func IAmAlive(context echo.Context) error {
	err := service.Test()

	if err != nil {
		err = errors.NewNotSupported(err, "error.Handler-get-error-from-test-service")
		logger.Error(err)
		return err
	}

	return context.JSON(http.StatusOK, "I Am Alive")
}
