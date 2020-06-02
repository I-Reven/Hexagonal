package socket

import socketio "github.com/graarh/golang-socketio"

type Default interface {
	Connect(c *socketio.Channel)
	Disconnect(c *socketio.Channel)
	Error(c *socketio.Channel)
}
