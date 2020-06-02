package core

import "github.com/gin-gonic/gin"

var (
	engine *gin.Engine
)

type Kernel struct {
	Middleware Middleware
	Worker     Worker
}

func (k Kernel) Boot() {
	k.Middleware.Handler()
	k.Worker.Work()
}
