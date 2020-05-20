package core

import (
	"github.com/I-Reven/Hexagonal/src/infrastructures/logger"
	"github.com/labstack/echo"
)

func Boot(e *echo.Echo) {
	Middleware(e)
	Route(e)
	Worker(e)
}

func BootDependencies(e *echo.Echo) {
	logger.Boot(e)
}
