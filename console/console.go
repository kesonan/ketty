package console

import (
	"fmt"
	"os"
	"runtime"
	"sync"
	"time"

	"github.com/anqiansong/ketty"
	"github.com/anqiansong/ketty/text"
	"github.com/logrusorgru/aurora"
	"github.com/pkg/errors"
)

const dateFormat = `2006-01-02 15:04:05.999`

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

type Console struct {
	useColor bool
	opt      []text.Option
	once     sync.Once
}

func NewConsole(opt ...text.Option) ketty.Logger {
	return &Console{
		opt:      opt,
		useColor: true,
	}
}

func (c *Console) DisableColor() {
	c.once.Do(func() {
		c.useColor = false
	})
}

func (c *Console) DisableBorder() {
	c.once.Do(func() {
		c.opt = append(c.opt, text.DisableBorder())
	})
}

func (c *Console) Info(format string, v ...interface{}) {
	msg := fmt.Sprintf(format, v...)
	opt := []text.Option{text.WithPrefix("[INFO] ", time.Now().Format(dateFormat))}
	opt = append(opt, c.opt...)
	output := text.Convert(msg, opt...)
	if runtime.GOOS != "windows" && c.useColor {
		output = aurora.Green(output).String()
	}
	fmt.Fprintf(os.Stdout, output)
}

func (c *Console) Debug(format string, v ...interface{}) {
	msg := fmt.Sprintf(format, v...)
	opt := []text.Option{text.WithPrefix("[DEBUG] ", time.Now().Format(dateFormat))}
	opt = append(opt, c.opt...)
	output := text.Convert(msg, opt...)
	if runtime.GOOS != "windows" && c.useColor {
		output = aurora.Blue(output).String()
	}
	fmt.Fprintf(os.Stdout, output)
}

func (c *Console) Warn(format string, v ...interface{}) {
	msg := fmt.Sprintf(format, v...)
	opt := []text.Option{text.WithPrefix("[WARN] ", time.Now().Format(dateFormat))}
	opt = append(opt, c.opt...)
	output := text.Convert(msg, opt...)
	if runtime.GOOS != "windows" && c.useColor {
		output = aurora.Yellow(output).String()
	}
	fmt.Fprintf(os.Stdout, output)
}

func (c *Console) Error(err error) {
	err = errors.WithStack(err)
	msg := fmt.Sprintf("%+v", err)
	opt := []text.Option{text.WithPrefix("[ERROR] ", time.Now().Format(dateFormat))}
	opt = append(opt, c.opt...)
	output := text.Convert(msg, opt...)
	if runtime.GOOS != "windows" && c.useColor {
		output = aurora.Red(output).String()
	}
	fmt.Fprintf(os.Stderr, output)
}
