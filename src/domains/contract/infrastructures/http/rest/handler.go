package rest

import "github.com/gin-gonic/gin"

type Handler interface {
	Handler(ctx *gin.Context)
}
