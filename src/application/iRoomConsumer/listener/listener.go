package listener

import (
	kernel "github.com/I-Reven/Hexagonal/src/application/iRoomConsumer"
	"github.com/I-Reven/Hexagonal/src/framework/listener"
	"github.com/I-Reven/Hexagonal/src/framework/logger"
	"net/http"
	"time"
)

type Listener struct {
	log    logger.Log
	listen listener.Listen
	http   kernel.Http
	socket kernel.Socket
	grpc   kernel.Grpc
}

func (l *Listener) Listen() {
	l.listen.Http(&http.Server{
		Addr:         ":83",
		Handler:      l.http.Route(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	})

	l.listen.Socket(&http.Server{
		Addr:         ":84",
		Handler:      l.socket.Route(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	})

	l.listen.Grpc("tcp", ":85", l.grpc.Route())

	l.listen.Run()
}
