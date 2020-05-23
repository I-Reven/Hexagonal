package main

import (
	kernel "github.com/I-Reven/Hexagonal/src/applications/core"
	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"os"
)

func init() {
	setEnv()
}

func main() {
	e := echo.New()

	kernel.BootDependencies(e)
	kernel.Boot(e)

	e.Logger.Fatal(e.Start(":80"))
}

func setEnv() {
	var err error
	env := os.Getenv("APP_ENV")

	switch env {
	case "production":
		err = godotenv.Load(".env")
	case "testing":
		err = godotenv.Load(".test.env")
	default:
		err = godotenv.Load(".local.env")
	}

	if err != nil {
		panic(err)
	}
}
