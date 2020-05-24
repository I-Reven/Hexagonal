package main

import (
	kernel "github.com/I-Reven/Hexagonal/src/applications/core"
	"github.com/I-Reven/Hexagonal/src/applications/core/console"
	"github.com/joho/godotenv"
	_ "github.com/mattn/go-colorable"
	"os"
)

func init() {
	setEnv()
	setOsArg()
}

func main() {
	boot()
	console.Cli()
}

func boot() {
	kernel.Boot()
}

func setOsArg() {
	if os.Getenv("PKG") == "" && len(os.Args) == 1 {
		os.Args = []string{os.Args[0], "serve"}
	}
}

func setEnv() {
	var err error
	env := os.Getenv("APP_ENV")

	switch env {
	case "production":
		err = godotenv.Load(".live.env")
	case "testing":
		err = godotenv.Load(".test.env")
	default:
		err = godotenv.Load(".local.env")
	}

	if err != nil {
		panic(err)
	}
}