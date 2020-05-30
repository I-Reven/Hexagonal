package core

import (
	"github.com/I-Reven/Hexagonal/src/infrastructures/http/rest"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Http struct{}

var (
	engine *gin.Engine
)

func init() {
	engine = gin.Default()
}

func (Http) Route() http.Handler {
	engine.GET("/ping", rest.Ping)
	engine.GET("/i-am-alive", rest.IAmAlive)
	engine.GET("/track/:trackId", rest.RequestTracker)

	return engine
}
