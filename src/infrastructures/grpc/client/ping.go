package client

import (
	"context"
	domain "github.com/I-Reven/Hexagonal/src/domains/grpc"
	"github.com/I-Reven/Hexagonal/src/infrastructures/logger"
	"github.com/juju/errors"
	"google.golang.org/grpc"
)

type Ping struct {
	log logger.Log
}

func (p Ping) Ping(addr string, request *domain.PingRequest) (*domain.PingResponse, error) {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())

	if err != nil {
		err = errors.NewNotSupported(err, "error.can-not-connect-to-grpc")
		p.log.Error(err)
		return nil, err
	}

	defer conn.Close()

	return domain.NewPingClient(conn).Ping(context.Background(), request)
}
