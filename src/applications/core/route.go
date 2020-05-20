package core

import (
	handler "github.com/I-Reven/Hexagonal/src/infrastructures/http/rest/core"
	"github.com/labstack/echo"
)

func Route(e *echo.Echo) {
	e.GET("/i-am-alive", handler.IAmAlive)
}
