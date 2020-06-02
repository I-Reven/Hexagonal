package applications

import "google.golang.org/grpc"

type Grpc interface {
	Route() *grpc.Server
}
