package iRoomProducer

import (
	"github.com/I-Reven/Hexagonal/src/application/iRoomProducer/service"
	"github.com/I-Reven/Hexagonal/src/domain/http/request"
	"github.com/I-Reven/Hexagonal/src/domain/http/response"
	"github.com/I-Reven/Hexagonal/src/framework/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Room struct {
	createRequest         request.CreateRoom
	addUserRequest        request.AddUser
	addMessageRequest     request.AddMessage
	seenMessageRequest    request.SeenMessage
	deliverMessageRequest request.DeliverMessage
	addMetaDataRequest    request.AddMetadata
	response              response.Response
	service               service.RoomProducer
	log                   logger.Log
}

func (h Room) Create(ctx *gin.Context) {
	if err := ctx.ShouldBindJSON(&h.createRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, h.response.BadRequest(err, "message.can-not-un-marshal-json"))
		return
	}

	if err := h.createRequest.Validate(); err != nil {
		ctx.JSON(http.StatusBadRequest, h.response.BadRequest(err, "message.invalid-json-data"))
		return
	}

	if err := h.service.Create(h.createRequest.CustomerName, h.createRequest.RoomId, h.createRequest.UserId); err != nil {
		ctx.JSON(http.StatusInternalServerError, h.response.TryAgain(err, "message.can-not-produce-create-room-message"))
		return
	}

	ctx.JSON(http.StatusAccepted, h.response.Success("OK"))
	return
}

func (h Room) AddUser(ctx *gin.Context) {
	if err := ctx.ShouldBindJSON(&h.addUserRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, h.response.BadRequest(err, "message.can-not-un-marshal-json"))
		return
	}

	if err := h.addUserRequest.Validate(); err != nil {
		ctx.JSON(http.StatusBadRequest, h.response.BadRequest(err, "message.invalid-json-data"))
		return
	}

	if err := h.service.AddUser(h.addUserRequest.CustomerName, h.addUserRequest.RoomId, h.addUserRequest.UserId); err != nil {
		ctx.JSON(http.StatusInternalServerError, h.response.TryAgain(err, "message.can-not-produce-add-user-message"))
		return
	}

	ctx.JSON(http.StatusAccepted, h.response.Success("OK"))
	return
}

func (h Room) AddMessage(ctx *gin.Context) {
	if err := ctx.ShouldBindJSON(&h.addMessageRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, h.response.BadRequest(err, "message.can-not-un-marshal-json"))
		return
	}

	if err := h.addMessageRequest.Validate(); err != nil {
		ctx.JSON(http.StatusBadRequest, h.response.BadRequest(err, "message.invalid-json-data"))
		return
	}

	if err := h.service.AddMessage(h.addMessageRequest.CustomerName, h.addMessageRequest.RoomId, h.addMessageRequest.UserId, h.addMessageRequest.Content, h.addMessageRequest.Kind); err != nil {
		ctx.JSON(http.StatusInternalServerError, h.response.TryAgain(err, "message.can-not-produce-add-message-message"))
		return
	}

	ctx.JSON(http.StatusAccepted, h.response.Success("OK"))
	return
}

func (h Room) SeenMessage(ctx *gin.Context) {
	if err := ctx.ShouldBindJSON(&h.seenMessageRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, h.response.BadRequest(err, "message.can-not-un-marshal-json"))
		return
	}

	if err := h.seenMessageRequest.Validate(); err != nil {
		ctx.JSON(http.StatusBadRequest, h.response.BadRequest(err, "message.invalid-json-data"))
		return
	}

	if err := h.service.SeenMessage(h.seenMessageRequest.CustomerName, h.seenMessageRequest.RoomId, h.seenMessageRequest.MessageId, h.seenMessageRequest.UserId); err != nil {
		ctx.JSON(http.StatusInternalServerError, h.response.TryAgain(err, "message.can-not-produce-seen-message-message"))
		return
	}

	ctx.JSON(http.StatusAccepted, h.response.Success("OK"))
	return
}

func (h Room) DeliverMessage(ctx *gin.Context) {
	if err := ctx.ShouldBindJSON(&h.deliverMessageRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, h.response.BadRequest(err, "message.can-not-un-marshal-json"))
		return
	}

	if err := h.deliverMessageRequest.Validate(); err != nil {
		ctx.JSON(http.StatusBadRequest, h.response.BadRequest(err, "message.invalid-json-data"))
		return
	}

	if err := h.service.DeliverMessage(h.deliverMessageRequest.CustomerName, h.deliverMessageRequest.RoomId, h.deliverMessageRequest.MessageId, h.deliverMessageRequest.UserId); err != nil {
		ctx.JSON(http.StatusInternalServerError, h.response.TryAgain(err, "message.can-not-produce-deliver-message-message"))
		return
	}

	ctx.JSON(http.StatusAccepted, h.response.Success("OK"))
	return
}

func (h Room) AddMetaData(ctx *gin.Context) {
	if err := ctx.ShouldBindJSON(&h.addMetaDataRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, h.response.BadRequest(err, "message.can-not-un-marshal-json"))
		return
	}

	if err := h.addMetaDataRequest.Validate(); err != nil {
		ctx.JSON(http.StatusBadRequest, h.response.BadRequest(err, "message.invalid-json-data"))
		return
	}

	if err := h.service.AddMetaData(h.addMetaDataRequest.CustomerName, h.addMetaDataRequest.RoomId, h.addMetaDataRequest.Key, h.addMetaDataRequest.Kind, h.addMetaDataRequest.Value); err != nil {
		ctx.JSON(http.StatusInternalServerError, h.response.TryAgain(err, "message.can-not-produce-deliver-message-message"))
		return
	}

	ctx.JSON(http.StatusAccepted, h.response.Success("OK"))
	return
}
