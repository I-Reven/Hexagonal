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

type Socket struct{}

func init() {
	serveMux = http.NewServeMux()
	socket = socketio.NewServer(transport.GetDefaultWebsocketTransport())
}

func (Socket) Route() *http.ServeMux {

	socket.On(socketio.OnConnection, handler.Connect)
	socket.On(socketio.OnDisconnection, handler.Disconnect)
	socket.On(socketio.OnError, handler.Error)

	serveMux.Handle("/io/", socket)
	return serveMux
}
