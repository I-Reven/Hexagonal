package logger

import (
	"github.com/I-Reven/Hexagonal/src/domains/entity"
	"github.com/gin-gonic/gin"
)

type Tracker interface {
	Create(context *gin.Context) (string, error)
	Message(message string) error
	Error(error error) error
	Data(data ...interface{}) error
	Debug(message string, data ...interface{}) error
	Get() (entity.Track, error)
}
