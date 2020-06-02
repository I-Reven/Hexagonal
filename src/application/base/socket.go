package core

import (
	socketio "github.com/graarh/golang-socketio"
	"github.com/graarh/golang-socketio/transport"
	"net/http"
)

var (
	serveMux *http.ServeMux
	socket   *socketio.Server
)

type Socket struct{}

func init() {
	serveMux = http.NewServeMux()
	socket = socketio.NewServer(transport.GetDefaultWebsocketTransport())
}

func (s Socket) Route() *http.ServeMux {

	serveMux.Handle("/io/", socket)
	return serveMux
}
