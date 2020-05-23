package socket

import (
	"github.com/I-Reven/Hexagonal/src/infrastructures/logger"
	socketio "github.com/graarh/golang-socketio"
)

//Connect Socket
func Connect(c *socketio.Channel) {
	logger.Info("Connect to socket " + c.Id())
}

//Disconnect Socket
func Disconnect(c *socketio.Channel) {
	logger.Info("Disconnect to socket " + c.Id())
}

//Error Socket
func Error(c *socketio.Channel) {
	logger.Info("Error to socket " + c.Id())
}
