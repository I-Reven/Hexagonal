package core

import (
	domain "github.com/I-Reven/Hexagonal/src/domains/grpc"
	"github.com/I-Reven/Hexagonal/src/infrastructures/grpc/handler"
	"google.golang.org/grpc"
)

type Grpc struct{}

func (Grpc) Route() *grpc.Server {
	server := grpc.NewServer()

	domain.RegisterPingServer(server, &handler.Ping{})

	return server
}
