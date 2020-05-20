package logger

type (
	Logger interface {
		Debug(i ...interface{})
		Info(i ...interface{})
		Warn(i ...interface{})
		Error(i ...interface{})
		Fatal(i ...interface{})
		Panic(i ...interface{})
	}
)
