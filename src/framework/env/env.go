package env

import (
	"github.com/joho/godotenv"
	"os"
)

type Env struct{}

func (e *Env) SetOsArg() {
	if os.Getenv("AUTO_SERVE") == "true" && len(os.Args) == 1 {
		os.Args = []string{os.Args[0], "serve"}
	}
}

func (e *Env) SetEnv() {
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
