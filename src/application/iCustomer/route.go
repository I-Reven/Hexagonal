package core

import (
	"github.com/I-Reven/Hexagonal/src/infrastructure/http/iCustomer"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Http struct {
	iCustomer.CreateCustomer
}

func init() {
	engine = gin.Default()
}

func (h Http) Route() http.Handler {
	v1 := engine.Group("/v1")
	{
		v1.POST("/create-customer", h.CreateCustomer.Handler)
	}

	return engine
}
