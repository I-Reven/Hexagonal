package core

import (
	domain "github.com/I-Reven/Hexagonal/src/domain/grpc"
	"github.com/I-Reven/Hexagonal/src/infrastructure/grpc/core/handler"
	"google.golang.org/grpc"
)

type Grpc struct{}

func (Grpc) Route() *grpc.Server {
	server := grpc.NewServer()

	domain.RegisterPingServer(server, &handler.Ping{})

	return server
}
