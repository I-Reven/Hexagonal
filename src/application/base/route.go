package core

import (
	"github.com/I-Reven/Hexagonal/src/infrastructure/http/core"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Http struct {
	iAmAlive core.IAmAlive
	tracker  core.Tracker
}

func init() {
	engine = gin.Default()
}

func (h Http) Route() http.Handler {
	_ = engine.Group("/v1")
	{

	}

	return engine
}
