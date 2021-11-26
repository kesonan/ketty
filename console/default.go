package console

func DisableColor() {
	colorConsole.DisableColor()
}

func DisableBorder() {
	colorConsole.DisableBorder()
}

var colorConsole = NewConsole()

func Info(format string, v ...interface{}) {
	colorConsole.Info(format, v...)
}

func Debug(format string, v ...interface{}) {
	colorConsole.Debug(format, v...)
}

func Warn(format string, v ...interface{}) {
	colorConsole.Warn(format, v...)
}

func Error(err error) {
	colorConsole.Error(err)
}
