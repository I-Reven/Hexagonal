package iRoom

import (
	"github.com/I-Reven/Hexagonal/src/application/iRoom/service"
	"github.com/I-Reven/Hexagonal/src/domain/http/request"
	"github.com/I-Reven/Hexagonal/src/domain/http/response"
	"github.com/I-Reven/Hexagonal/src/framework/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Room struct {
	request request.CreateRoom

	response response.Response
	service service.Room
	log logger.Log
}

func (h Room)Create(ctx *gin.Context)  {
	if err := ctx.ShouldBindJSON(&h.request); err != nil {
		ctx.JSON(http.StatusBadRequest, h.response.BadRequest(err, "message.can-not-un-marshal-json"))
		return
	}

	if err := h.request.Validate(); err != nil {
		ctx.JSON(http.StatusBadRequest, h.response.BadRequest(err, "message.invalid-json-data"))
		return
	}

	if err := h.service.Create(h.request.CustomerName, h.request.RoomId, h.request.UserId); err != nil {
		ctx.JSON(http.StatusInternalServerError, h.response.TryAgain(err, "message.can-not-produce-message-error-try-again"))
	}

	ctx.JSON(http.StatusCreated, h.response.Success("OK"))

}