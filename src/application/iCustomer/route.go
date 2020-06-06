package core

import (
	"github.com/I-Reven/Hexagonal/src/infrastructure/http/iCustomer"
	createCustomerWebHook "github.com/I-Reven/Hexagonal/src/infrastructure/webHook/iCustomer"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Http struct {
	createCustomer        iCustomer.CreateCustomer
	createCustomerWebHook createCustomerWebHook.CreateCustomer
}

func init() {
	engine = gin.Default()
}

func (h Http) Route() http.Handler {
	webHook := engine.Group("/web-hook")
	{
		webHook.GET("/create-customer/:token", h.createCustomerWebHook.Approve)
		webHook.GET("/create-customer-cancel/:token", h.createCustomerWebHook.Cancel)
	}

	v1 := engine.Group("/v1")
	{
		v1.POST("/create-customer", h.createCustomer.Handler)
	}

	return engine
}
