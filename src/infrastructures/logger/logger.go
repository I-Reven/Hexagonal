package logger

import (
	"github.com/gin-gonic/gin"
	"github.com/labstack/gommon/log"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"time"
)

func init() {
	logrus.SetFormatter(&logrus.TextFormatter{})
	logrus.SetReportCaller(false)
	logrus.SetFormatter(&logrus.TextFormatter{
		DisableColors: false,
		FullTimestamp: true,
	})
}

//Boot Logger
func Boot() {
	setLoggerOutput()
}

//Set Logger Output
func setLoggerOutput() {
	if os.Getenv("APP_DEBUG") == "true" {
		log.SetOutput(os.Stderr)
		log.SetLevel(log.DEBUG)
	} else {
		currentTime := time.Now()
		logPath, err := os.OpenFile("/var/log/hex/"+currentTime.Format("2006-01-02-")+os.Getenv("PKG")+".log", os.O_APPEND|os.O_CREATE, 0755)

		if err != nil {
			log.Fatal(err)
		}

		gin.DisableConsoleColor()
		gin.DefaultWriter = io.MultiWriter(logPath)

		//log.SetOutput(logPath)
		//log.SetLevel(log.WARN)
		//log.SetLevel(log.DEBUG)
	}

}

//Debug log
func Debug(i ...interface{}) {
	logrus.Debug(i...)
	log.Debug(i...)
}

//Info log
func Info(i ...interface{}) {
	logrus.Info(i...)
	log.Info(i...)
}

//Warn log
func Warn(i ...interface{}) {
	logrus.Warn(i...)
	log.Warn(i...)
}

//Error log
func Error(i ...interface{}) {
	logrus.Error(i...)
	log.Error(i...)
}

//Fatal log
func Fatal(i ...interface{}) {
	logrus.Fatal(i...)
	log.Fatal(i...)
}

//Panic log
func Panic(i ...interface{}) {
	logrus.Panic(i...)
	log.Panic(i...)
}
