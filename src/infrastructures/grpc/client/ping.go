package client

import (
	"context"
	domain "github.com/I-Reven/Hexagonal/src/domains/grpc"
	"github.com/juju/errors"
	"google.golang.org/grpc"
)

type Ping struct{}

func (Ping) Ping(request *domain.PingRequest, port string) (*domain.PingResponse, error) {
	conn, err := grpc.Dial(port, grpc.WithInsecure())

	if err != nil {
		errors.NewNotSupported(err, "error.can-not-connect-to-grpc")
		return nil, err
	}

	defer conn.Close()

	return domain.NewPingClient(conn).Ping(context.Background(), request)
}
