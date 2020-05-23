package logger

import (
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"github.com/sirupsen/logrus"
	"os"
	"sync"
	"time"
)

var (
	once sync.Once
	Echo *echo.Echo
)

func init() {
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
		SetLoggerOutput()
	})
}

func SetLoggerOutput() {
	if os.Getenv("APP_DEBUG") == "true" {
		log.SetOutput(os.Stderr)
		log.SetLevel(log.DEBUG)
	} else {
		currentTime := time.Now()
		logPath, err := os.OpenFile("/var/log/hex/"+currentTime.Format("2006-01-02-")+os.Getenv("PKG")+".log", os.O_APPEND|os.O_CREATE, 0755)

		if err != nil {
			Echo.Logger.Fatal(err)
		}

		log.SetOutput(logPath)
		log.SetLevel(log.WARN)
	}
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
