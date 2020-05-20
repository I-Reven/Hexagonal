package core

import (
	"github.com/I-Reven/Hexagonal/applications/core/service"
	"github.com/labstack/echo"
	"net/http"
)

func IAmAlive(context echo.Context) error {
	err := service.Test()

	if err != nil {
		return err
	}

	return context.JSON(http.StatusOK, "I Am Alive")
}
