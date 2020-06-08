package core

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Http struct{}

func init() {
	engine = gin.Default()
}

func (h Http) Route() http.Handler {
	_ = engine.Group("/v1")
	{
	}

	return engine
}
