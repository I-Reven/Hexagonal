package core

import (
	"github.com/I-Reven/Hexagonal/src/infrastructures/http/rest"
	"net/http"
)

func Route() http.Handler {

	engine.GET("/ping", rest.Ping)
	engine.GET("/i-am-alive", rest.IAmAlive)
	engine.GET("/track/:trackId", rest.RequestTracker)

	return engine
}
