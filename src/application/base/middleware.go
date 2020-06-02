package core

import (
	"github.com/I-Reven/Hexagonal/src/application/core/middleware"
	redis "github.com/I-Reven/Hexagonal/src/infrastructure/repository/redis/session"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type Middleware struct {
	Session redis.Session
	Tracker middleware.Tracker
}

func (m Middleware) Handler() {
	engine.Use(gin.Logger())
	engine.Use(gin.Recovery())

	engine.Use(sessions.Sessions("iCustomer", m.Session.Store()))
	engine.Use(m.Tracker.Handler())
}
