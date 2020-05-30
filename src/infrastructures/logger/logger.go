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
