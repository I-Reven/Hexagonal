package core

import (
	"google.golang.org/grpc"
)

type Grpc struct{}

func (Grpc) Route() *grpc.Server {
	server := grpc.NewServer()

	return server
}
