package main

import (
	kernel "github.com/I-Reven/Hexagonal/src/applications/core"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	kernel.BootDependencies(e)
	kernel.Boot(e)

	e.Logger.Fatal(e.Start(":80"))
}
