package iCustomer

import (
	"github.com/I-Reven/Hexagonal/src/application/iCustomer/service"
	"github.com/I-Reven/Hexagonal/src/framework/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CreateCustomer struct {
	log     logger.Log
	service service.Customer
}

func (h *CreateCustomer) Handler(ctx *gin.Context) {

	ctx.JSON(http.StatusOK, "CreateCustomer")
}
