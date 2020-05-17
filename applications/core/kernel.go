package core

import (
	handler "github.com/I-Reven/Hexagonal/infrastructures/http/rest/core"
	"github.com/labstack/echo"
)

func Boot(e *echo.Echo) {
	e.GET("/i-am-alive", handler.IAmAlive)
}
