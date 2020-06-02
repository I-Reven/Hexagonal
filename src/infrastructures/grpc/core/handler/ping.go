package handler

import (
	"context"
	domain "github.com/I-Reven/Hexagonal/src/domains/grpc"
	"github.com/I-Reven/Hexagonal/src/infrastructures/logger"
)

type Ping struct {
	log logger.Log
}

func (p Ping) Ping(ctx context.Context, in *domain.PingRequest) (*domain.PingResponse, error) {
	p.log.TraceLn("Receive message %s", in.Message)
	return &domain.PingResponse{Message: "PONG"}, nil
}
