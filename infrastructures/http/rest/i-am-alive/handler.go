package i_am_alive

import (
	"github.com/labstack/echo"
	"net/http"
)

func IAmAlive(context echo.Context) {
	return context.JSON(http.StatusOK, )
}
