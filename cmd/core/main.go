package main

import (
	kernel "github.com/I-Reven/Hexagonal/applications/core"
	"github.com/I-Reven/Hexagonal/infrastructures/logger"
	"github.com/labstack/echo"
)

func main ()  {
	e := echo.New()

	logger.Boot(e)
	kernel.BootDependencies(e)
	kernel.Boot(e)

	logger.LOG().FatalE(e.Start(":80"))
}

