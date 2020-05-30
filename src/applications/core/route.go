package core

import (
	"github.com/I-Reven/Hexagonal/src/infrastructures/http/rest"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Http struct {
	iAmAlive rest.IAmAlive
	ping     rest.Ping
	tracker  rest.Tracker
}

var (
	engine *gin.Engine
)

func init() {
	engine = gin.Default()
}

func (h Http) Route() http.Handler {

	engine.GET("/ping", h.ping.Handler)
	engine.GET("/i-am-alive", h.iAmAlive.Handler)
	engine.GET("/track/:trackId", h.tracker.Handler)

	return engine
}
