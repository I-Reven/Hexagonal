package applications

import "net/http"

type Route interface {
	Route() http.Handler
}
