package applications

import "net/http"

type Socket interface {
	Route() *http.ServeMux
}
