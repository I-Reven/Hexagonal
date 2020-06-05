package iCustomer

import (
	"github.com/I-Reven/Hexagonal/src/application/iCustomer/service"
	"github.com/I-Reven/Hexagonal/src/domain/http/request"
	"github.com/I-Reven/Hexagonal/src/domain/http/response"
	"github.com/I-Reven/Hexagonal/src/framework/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CreateCustomer struct {
	log      logger.Log
	service  service.Customer
	request  request.CreateCustomer
	response response.Response
}

func (h *CreateCustomer) Handler(ctx *gin.Context) {

	if err := ctx.ShouldBindJSON(&h.request); err != nil {
		ctx.JSON(http.StatusBadRequest, h.response.BadRequest(err, "message.can-not-un-marshal-json"))
		return
	}

	if err := h.request.Validate(); err != nil {
		ctx.JSON(http.StatusBadRequest, h.response.BadRequest(err, "message.invalid-json-data"))
		return
	}

	if err := h.service.CreateProducer(h.request.CustomerName); err != nil {
		ctx.JSON(http.StatusInternalServerError, h.response.TryAgain(err, "message.can-not-produce-message-error-try-again"))
	}

	ctx.JSON(http.StatusCreated, h.response.Success("OK"))
}
