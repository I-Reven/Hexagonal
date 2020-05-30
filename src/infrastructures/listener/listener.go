package listener

import (
	"github.com/I-Reven/Hexagonal/src/infrastructures/logger"
	"github.com/juju/errors"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"net"
	"net/http"
)

type Listen struct {
	group errgroup.Group
	log   logger.Log
}

func (l *Listen) Http(server *http.Server) {
	l.group.Go(func() error {
		return server.ListenAndServe()
	})
}

func (l *Listen) Socket(server *http.Server) {
	l.group.Go(func() error {
		return server.ListenAndServe()
	})
}

func (l *Listen) Grpc(network, address string, server *grpc.Server) {
	listener, err := net.Listen(network, address)
	l.group.Go(func() error {

		if err != nil {
			return errors.NewNotSupported(err, "Server can not listen to  GRPC")
		}

		return server.Serve(listener)
	})
}

func (l *Listen) Run() {
	if err := l.group.Wait(); err != nil {
		l.log.Fatal(err)
	}
}
