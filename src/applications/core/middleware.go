package core

import (
	core "github.com/I-Reven/Hexagonal/src/applications/core/middleware"
	"github.com/I-Reven/Hexagonal/src/infrastructures/repository/redis/session"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"time"
)

func middleware() {
	engine.Use(gin.Logger())
	engine.Use(gin.Recovery())
	engine.Use(sessions.Sessions("core", session.Store()))
	engine.Use(core.RequestTracker())
	engine.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "DELETE", "PUT", "PATCH"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "*"
		},
		MaxAge: 12 * time.Hour,
	}))
}
