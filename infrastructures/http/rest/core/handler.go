package core

import (
	"github.com/I-Reven/Hexagonal/applications/core/service"
	"github.com/labstack/echo"
	"net/http"
)

func IAmAlive(context echo.Context) error {
	err := service.TestDatabase()

	if err != nil {
		return err
	}

	return context.String(http.StatusOK, "I Am Alive")
}
