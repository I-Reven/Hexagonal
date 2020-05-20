package logger

import (
	"github.com/labstack/gommon/log"
)

type (
	Logger interface {
		Debug(i ...interface{})
		DebugJ(j log.JSON)
		DebugE(e error)
		Info(i ...interface{})
		InfoJ(j log.JSON)
		InfoE(e error)
		Warn(i ...interface{})
		WarnJ(j log.JSON)
		WarnE(e error)
		Error(i ...interface{})
		ErrorJ(j log.JSON)
		ErrorE(e error)
		Fatal(i ...interface{})
		FatalJ(j log.JSON)
		FatalE(e error)
		Panic(i ...interface{})
		PanicJ(j log.JSON)
		PanicE(e error)
	}
)
