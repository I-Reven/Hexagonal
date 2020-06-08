package core

import (
	"github.com/I-Reven/Hexagonal/src/infrastructure/http/iRoomProducer"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Http struct {
	iRoom iRoomProducer.Room
}

func init() {
	engine = gin.Default()
}

func (h Http) Route() http.Handler {
	v1 := engine.Group("/v1")
	{
		v1.POST("/room/create", h.iRoom.Create)
		v1.PATCH("/room/add-user", h.iRoom.AddUser)
		v1.PATCH("/room/add-message", h.iRoom.AddMessage)
		v1.PATCH("/room/add-meta-data", h.iRoom.AddMetaData)
		v1.PUT("/room/seen-message", h.iRoom.SeenMessage)
		v1.PUT("/room/deliver-message", h.iRoom.DeliverMessage)
	}

	return engine
}
