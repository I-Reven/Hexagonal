package core

import (
	"github.com/gin-gonic/gin"
	socketio "github.com/graarh/golang-socketio"
	"github.com/graarh/golang-socketio/transport"
	"net/http"
)

var (
	engine   *gin.Engine
	serveMux *http.ServeMux
	socket   *socketio.Server
)

func init() {
	engine = gin.Default()
	serveMux = http.NewServeMux()
	socket = socketio.NewServer(transport.GetDefaultWebsocketTransport())
}

func Boot() {
	middleware()
	worker()
}
