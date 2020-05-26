package logger

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"io"
	"log"
	"os"
	"time"
)

func init() {
	SetLogPath()
	logrus.SetFormatter(&logrus.TextFormatter{})
	logrus.SetReportCaller(false)
	logrus.SetFormatter(&logrus.TextFormatter{
		DisableColors: false,
		FullTimestamp: true,
	})
	logrus.SetLevel(logrus.WarnLevel)
}

func SetLogPath() {
	if os.Getenv("PKG") == "" {
		log.SetOutput(os.Stderr)
	} else {
		currentTime := time.Now()
		logPath, err := os.Create("/var/log/hex/" + currentTime.Format("2006-01-02-") + os.Getenv("PKG") + ".log")
		logPathGin, err := os.Create("/var/log/hex/" + currentTime.Format("2006-01-02-") + os.Getenv("PKG") + ".gin.log")
		logPathLogrus, err := os.Create("/var/log/hex/" + currentTime.Format("2006-01-02-") + os.Getenv("PKG") + ".logrus.log")

		if err != nil {
			log.Fatal(err)
		}

		gin.DisableConsoleColor()
		gin.DefaultWriter = io.MultiWriter(logPathGin)
		log.SetOutput(io.MultiWriter(logPath))
		logrus.SetOutput(io.MultiWriter(logPathLogrus))
	}
}

func Debug(message string, data ...interface{}) {
	if os.Getenv("APP_DEBUG") == "true" {
		TrackDebug(message, data...)
		logrus.Debug(message)
		logrus.Debug(data...)
	}
}

func TraceLn(data ...interface{}) {
	TrackData(data...)
	logrus.Traceln(data...)
}

func Info(message string) {
	TrackMessage(message)
	logrus.Info(message)
}

func Warn(err error) {
	TrackError(err)
	logrus.Warn(err)
}

func Error(err error) {
	TrackError(err)
	logrus.Error(err)
}

func Fatal(err error) {
	TrackError(err)
	logrus.Fatal(err)
	log.Fatal(err)
}

func Panic(err error) {
	TrackError(err)
	logrus.Panic(err)
	log.Panic(err)
}

func StartDebug() error {
	return os.Setenv("APP_DEBUG", "true")
}

func EndDebug() error {
	return os.Setenv("APP_DEBUG", "false")
}
