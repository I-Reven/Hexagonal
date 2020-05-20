package logger

import (
	"github.com/fatih/structs"
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
	"os"
	"sync"
)

var (
	once sync.Once
	Echo *echo.Echo
)

func init() {
	logrus.SetOutput(os.Stdout)
	logrus.SetFormatter(&logrus.TextFormatter{})
	logrus.SetReportCaller(false)
	logrus.SetFormatter(&logrus.TextFormatter{
		DisableColors: false,
		FullTimestamp: true,
	})
}

func Boot(e *echo.Echo) {
	once.Do(func() {
		Echo = e
	})
}

func Debug(i ...interface{}) {
	logrus.Debug(i...)
	Echo.Logger.Debug(i...)
}

func Info(i ...interface{}) {
	logrus.Info(i...)
	Echo.Logger.Info(i...)
}

func Warn(i ...interface{}) {
	logrus.Warn(i...)
	Echo.Logger.Warn(i...)
}

func Error(i ...interface{}) {
	logrus.Error(i...)
	Echo.Logger.Error(i...)
}

func Fatal(i ...interface{}) {
	logrus.Fatal(i...)
	Echo.Logger.Fatal(i...)
}

func Panic(i ...interface{}) {
	logrus.Panic(i...)
	Echo.Logger.Panic(i...)
}

func DebugF(atr interface{}, err error) {
	logrus.WithFields(structs.Map(atr)).Debug(err)
	Echo.Logger.Debug(err)
}

func InfoF(i ...interface{}) {
	logrus.Info(i...)
	Echo.Logger.Info(i...)
}

func WarnF(i ...interface{}) {
	logrus.Warn(i...)
	Echo.Logger.Warn(i...)
}

func ErrorF(i ...interface{}) {
	logrus.Error(i...)
	Echo.Logger.Error(i...)
}

func FatalF(i ...interface{}) {
	logrus.Fatal(i...)
	Echo.Logger.Fatal(i...)
}

func PanicF(i ...interface{}) {
	logrus.Panic(i...)
	Echo.Logger.Panic(i...)
}
