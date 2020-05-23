package core

import (
	"github.com/I-Reven/Hexagonal/src/infrastructures/http/rest"
	"github.com/labstack/echo"
)

func Route(e *echo.Echo) {
	e.GET("/i-am-alive", rest.IAmAlive)
}
