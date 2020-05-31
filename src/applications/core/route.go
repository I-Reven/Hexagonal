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
	v1 := engine.Group("/v1")
	{
		v1.GET("/ping", h.ping.Handler)
		v1.GET("/i-am-alive", h.iAmAlive.Handler)
		v1.POST("/track", h.tracker.Handler)
	}

	return engine
}
