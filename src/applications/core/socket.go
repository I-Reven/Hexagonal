package core

import (
	handler "github.com/I-Reven/Hexagonal/src/infrastructures/socket"
	socketio "github.com/graarh/golang-socketio"
	"github.com/graarh/golang-socketio/transport"
	"net/http"
)

var (
	serveMux *http.ServeMux
	socket   *socketio.Server
)

type Socket struct {
	socket handler.Default
}

func init() {
	serveMux = http.NewServeMux()
	socket = socketio.NewServer(transport.GetDefaultWebsocketTransport())
}

func (s Socket) Route() *http.ServeMux {

	socket.On(socketio.OnConnection, s.socket.Connect)
	socket.On(socketio.OnDisconnection, s.socket.Disconnect)
	socket.On(socketio.OnError, s.socket.Error)

	serveMux.Handle("/io/", socket)
	return serveMux
}
