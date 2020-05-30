package logger

import (
	"github.com/sirupsen/logrus"
	"log"
	"os"
)

type Log struct {
	tracker Tracker
}

func (l *Log) Debug(message string, data ...interface{}) {
	if os.Getenv("APP_DEBUG") == "true" {
		_ = l.tracker.Debug(message, data...)
		logrus.Debug(message)
		logrus.Debug(data...)
	}
}

func (l *Log) TraceLn(data ...interface{}) {
	_ = l.tracker.Data(data...)
	logrus.Traceln(data...)
}

func (l *Log) Info(message string) {
	_ = l.tracker.Message(message)
	logrus.Info(message)
}

func (l *Log) Warn(err error) {
	_ = l.tracker.Error(err)
	logrus.Warn(err)
}

func (l *Log) Error(err error) {
	_ = l.tracker.Error(err)
	logrus.Error(err)
}

func (l *Log) Fatal(err error) {
	_ = l.tracker.Error(err)
	logrus.Fatal(err)
	log.Fatal(err)
}

func (l *Log) Panic(err error) {
	_ = l.tracker.Error(err)
	logrus.Panic(err)
	log.Panic(err)
}

func (*Log) StartDebug() error {
	return os.Setenv("APP_DEBUG", "true")
}

func (*Log) EndDebug() error {
	return os.Setenv("APP_DEBUG", "false")
}
