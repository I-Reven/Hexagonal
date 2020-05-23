package main

import (
	kernel "github.com/I-Reven/Hexagonal/src/applications/core"
	"github.com/I-Reven/Hexagonal/src/infrastructures/logger"
	"github.com/joho/godotenv"
	"golang.org/x/sync/errgroup"
	"net/http"
	"os"
	"time"
)

var (
	g errgroup.Group
)

func init() {
	setEnv()
	boot()
}

func main() {
	engine := kernel.Route()
	socket := kernel.Socket()

	server := &http.Server{
		Addr:         ":80",
		Handler:      engine,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	serverMux := &http.Server{
		Addr:         ":81",
		Handler:      socket,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	g.Go(func() error {
		return server.ListenAndServe()
	})

	g.Go(func() error {
		return serverMux.ListenAndServe()
	})

	if err := g.Wait(); err != nil {
		logger.Fatal(err)
	}
}

func boot() {
	logger.Boot()
	kernel.Boot()
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
