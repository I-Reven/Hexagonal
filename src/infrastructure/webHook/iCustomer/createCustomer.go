package iCustomer

import (
	"github.com/I-Reven/Hexagonal/src/application/iCustomer/job"
	"github.com/I-Reven/Hexagonal/src/domain/http/response"
	"github.com/I-Reven/Hexagonal/src/infrastructure/repository/redis/webHook"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CreateCustomer struct {
	job        job.CreateCustomer
	repository webHook.WebHook
	response   response.Response
}

func (h *CreateCustomer) Approve(ctx *gin.Context) {
	token := ctx.Param("token")
	data, err := h.repository.Get("CreateCustomer", token)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, h.response.BadRequest(err, "Token invalid"))
		return
	}

	err, j := h.job.Init(data)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, h.response.BadRequest(err, "Can not Init Job"))
		return
	}

	err = j.Handler()

	if err != nil {
		j.Failed(err)
		ctx.JSON(http.StatusInternalServerError, h.response.InternalError(err, "Job error"))
		return
	}

	j.Done()
	ctx.JSON(http.StatusOK, h.response.Success("Ok"))
}

func (h *CreateCustomer) Cancel(ctx *gin.Context) {
	token := ctx.Param("token")

	if err := h.repository.Delete("CreateCustomer", token); err != nil {
		ctx.JSON(http.StatusBadRequest, h.response.BadRequest(err, "Token invalid"))
	}

	ctx.JSON(http.StatusOK, h.response.Success("Ok"))
}
