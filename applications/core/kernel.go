package core

import (
	"github.com/labstack/echo"
)

func Boot(e *echo.Echo) {
	Route(e)
	Worker()
}

func BootDependencies(e *echo.Echo) {

}
