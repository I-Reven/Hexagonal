package core

import (
	"github.com/I-Reven/Hexagonal/src/framework/logger"
	socketio "github.com/graarh/golang-socketio"
)

type Default struct {
	log logger.Log
}

func (h Default) Connect(c *socketio.Channel) {
	h.log.Info("Connect to socket " + c.Id())
}

func (h Default) Disconnect(c *socketio.Channel) {
	h.log.Info("Disconnect to socket " + c.Id())
}

func (h Default) Error(c *socketio.Channel) {
	h.log.Info("Error to socket " + c.Id())
}
