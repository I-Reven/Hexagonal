package core

import (
	"github.com/gin-gonic/gin"
)

//Middleware Http
func middleware() {
	engine.Use(gin.Logger())
	engine.Use(gin.Recovery())
}
