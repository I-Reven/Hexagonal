package handler

import (
	"context"
	domain "github.com/I-Reven/Hexagonal/src/domains/grpc"
	"log"
)

type Ping struct{}

func (Ping) Ping(ctx context.Context, in *domain.PingRequest) (*domain.PingResponse, error) {
	log.Printf("Receive message %s", in.Message)
	return &domain.PingResponse{Message: "PONG"}, nil
}
