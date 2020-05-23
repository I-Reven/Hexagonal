package core

import (
	"github.com/I-Reven/Hexagonal/src/infrastructures/http/rest"
	"net/http"
)

//Route Http
func Route() http.Handler {

	engine.GET("/i-am-alive", rest.IAmAlive)

	return engine
}
