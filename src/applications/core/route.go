package core

import (
	"github.com/I-Reven/Hexagonal/src/infrastructures/http/rest"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Http struct {
	iAmAlive rest.IAmAlive
	tracker  rest.Tracker
}

func init() {
	engine = gin.Default()
}

func (h Http) Route() http.Handler {
	v1 := engine.Group("/v1")
	{
		v1.GET("/i-am-alive", h.iAmAlive.Handler)
		v1.POST("/track", h.tracker.Handler)
	}

	return engine
}
