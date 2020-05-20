package logger

import (
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"sync"
)

var (
	once   sync.Once
	logger Logger
)

type (
	Logger struct {
		*echo.Echo
	}
)

func Boot(e *echo.Echo) {
	once.Do(func() {
		logger = Logger{e}
	})
}

func LOG() *Logger {
	return &logger
}

func (logger *Logger) Debug(i ...interface{}) {
	logger.Logger.Debug(i)
}

func (logger *Logger) DebugJ(j log.JSON) {
	logger.Logger.Debugj(j)
}

func (logger *Logger) DebugE(e error) {
	if e != nil {
		logger.Logger.Debug(e)
	}
}

func (logger *Logger) Info(i ...interface{}) {
	logger.Logger.Info(i)
}

func (logger *Logger) InfoJ(j log.JSON) {
	logger.Logger.Infoj(j)
}

func (logger *Logger) InfoE(e error) {
	if e != nil {
		logger.Logger.Info(e)
	}
}

func (logger *Logger) Warn(i ...interface{}) {
	logger.Logger.Warn(i)
}

func (logger *Logger) WarnJ(j log.JSON) {
	logger.Logger.Warnj(j)
}

func (logger *Logger) WarnE(e error) {
	if e != nil {
		logger.Logger.Warn(e)
	}
}

func (logger *Logger) Error(i ...interface{}) {
	logger.Logger.Error(i)
}

func (logger *Logger) ErrorJ(j log.JSON) {
	logger.Logger.Errorj(j)
}

func (logger *Logger) ErrorE(e error) {
	if e != nil {
		logger.Logger.Error(e)
	}
}

func (logger *Logger) Fatal(i ...interface{}) {
	logger.Logger.Fatal(i)
}

func (logger *Logger) FatalJ(j log.JSON) {
	logger.Logger.Fatalj(j)
}

func (logger *Logger) FatalE(e error) {
	if e != nil {
		logger.Logger.Fatal(e)
	}
}

func (logger *Logger) Panic(i ...interface{}) {
	logger.Logger.Panic(i)
}

func (logger *Logger) PanicJ(j log.JSON) {
	logger.Logger.Panicj(j)
}

func (logger *Logger) PanicE(e error) {
	if e != nil {
		logger.Logger.Panic(e)
	}
}
