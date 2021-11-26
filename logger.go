package ketty

type Logger interface {
	Info(format string, v ...interface{})
	Debug(format string, v ...interface{})
	Warn(format string, v ...interface{})
	Error(err error)
	Rotate() error
}
