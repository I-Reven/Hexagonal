package server

import (
	kernel "github.com/I-Reven/Hexagonal/src/applications/core"
	"github.com/I-Reven/Hexagonal/src/infrastructures/logger"
	"golang.org/x/sync/errgroup"
	"net/http"
	"time"
)

var (
	g errgroup.Group
)

// Listen
func Listen() {
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
