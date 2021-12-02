// MIT License
//
// Copyright (c) 2021 anqiansong
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package console

// DisableColor sets the color flag as false, if it is,
// the Console won't print log with color.
func DisableColor() {
	colorConsole.DisableColor()
}

// DisablePrefix prints log without prefix
func DisablePrefix() {
	colorConsole.DisablePrefix()
}

func SetOutput(dir string) {
	colorConsole.Use(WithOutputDir(dir))
}

// UsePrefix add a prefix to logger.
func UsePrefix(prefix string) {
	colorConsole.prefix = prefix
}

// Close for cleaning up.
func Close() {
	colorConsole.Close()
}

// DisableBorder prints log with borderless.
func DisableBorder() {
	colorConsole.DisableBorder()
}

// GetConsole returns a Console instance.
func GetConsole() *Console {
	return colorConsole
}

var colorConsole = NewConsole()

// Info prints info level log.
func Info(format string, v ...interface{}) {
	colorConsole.Info(format, v...)
}

// Debug prints debug level log.
func Debug(format string, v ...interface{}) {
	colorConsole.Debug(format, v...)
}

// Warn prints warn level log.
func Warn(format string, v ...interface{}) {
	colorConsole.Warn(format, v...)
}

// Error prints error level log.
func Error(err error) {
	colorConsole.Error(err)
}

// ErrorText prints error level log.
func ErrorText(format string, v ...interface{}) {
	colorConsole.ErrorText(format, v...)
}
