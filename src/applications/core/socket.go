package core

import (
	handler "github.com/I-Reven/Hexagonal/src/infrastructures/socket"
	socketio "github.com/graarh/golang-socketio"
	"net/http"
)

func Socket() *http.ServeMux {

	socket.On(socketio.OnConnection, handler.Connect)
	socket.On(socketio.OnDisconnection, handler.Disconnect)
	socket.On(socketio.OnError, handler.Error)

	serveMux.Handle("/io/", socket)
	return serveMux
}
