package main

import (
	kernel "github.com/I-Reven/Core/applications/i-am-alive"
	"github.com/labstack/echo"
)

func main ()  {
	e := echo.New()
	kernel.Boot(e)
	e.Logger.Fatal(e.Start(":80"))
}

