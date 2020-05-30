package logger

type Log interface {
	Debug(message string, data ...interface{})
	TraceLn(data ...interface{})
	Info(message string)
	Warn(err error)
	Error(err error)
	Fatal(err error)
	Panic(err error)
	StartDebug() error
	EndDebug() error
}
