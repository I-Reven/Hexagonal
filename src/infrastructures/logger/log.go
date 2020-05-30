package logger

import (
	"github.com/sirupsen/logrus"
	"log"
	"os"
)

type Log struct {
	Track Tracker
}

func (l Log) Debug(message string, data ...interface{}) {
	if os.Getenv("APP_DEBUG") == "true" {
		_ = l.Track.Debug(message, data...)
		logrus.Debug(message)
		logrus.Debug(data...)
	}
}

func (l Log) TraceLn(data ...interface{}) {
	_ = l.Track.Data(data...)
	logrus.Traceln(data...)
}

func (l Log) Info(message string) {
	_ = l.Track.Message(message)
	logrus.Info(message)
}

func (l Log) Warn(err error) {
	_ = l.Track.Error(err)
	logrus.Warn(err)
}

func (l Log) Error(err error) {
	_ = l.Track.Error(err)
	logrus.Error(err)
}

func (l Log) Fatal(err error) {
	_ = l.Track.Error(err)
	logrus.Fatal(err)
	log.Fatal(err)
}

func (l Log) Panic(err error) {
	_ = l.Track.Error(err)
	logrus.Panic(err)
	log.Panic(err)
}

func (Log) StartDebug() error {
	return os.Setenv("APP_DEBUG", "true")
}

func (Log) EndDebug() error {
	return os.Setenv("APP_DEBUG", "false")
}
