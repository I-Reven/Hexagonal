package core

import (
	"github.com/I-Reven/Hexagonal/src/applications/core/middleware"
	redis "github.com/I-Reven/Hexagonal/src/infrastructure/repository/redis/session"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"time"
)

type Middleware struct {
	Session redis.Session
	Tracker middleware.Tracker
}

func (m Middleware) Handler() {
	engine.Use(gin.Logger())
	engine.Use(gin.Recovery())

	engine.Use(sessions.Sessions("core", m.Session.Store()))
	engine.Use(m.Tracker.Handler())

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
