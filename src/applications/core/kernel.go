package core

import (
	"github.com/I-Reven/Hexagonal/src/infrastructures/logger"
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
	logger.SetLogPath()

	engine = gin.Default()
	serveMux = http.NewServeMux()
	socket = socketio.NewServer(transport.GetDefaultWebsocketTransport())
}

func Boot() {
	logger.Boot()
	middleware()
	worker()
}
