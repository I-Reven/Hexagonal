package rest

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Ping struct{}

func (h *Ping) Handler(ctx *gin.Context) {
	ctx.String(http.StatusOK, "PONG")
}
