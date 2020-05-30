package listener

import (
	"google.golang.org/grpc"
	"net/http"
)

type Listener interface {
	Http(server *http.Server)
	Socket(server *http.Server)
	Grpc(network, address string, server *grpc.Server)
	Run()
}
