package console

import (
	"fmt"
	"os"
	"runtime"
	"time"

	"github.com/anqiansong/ketty"
	"github.com/anqiansong/ketty/text"
	"github.com/logrusorgru/aurora"
	"github.com/pkg/errors"
)

const dateFormat = `2006-01-02 15:04:05.999`

type Option func(c *Console)

var _ ketty.Logger = (*Console)(nil)

type Console struct {
	useColor bool
	opt      []text.Option
	output   string
}

func NewConsole(opt ...Option) *Console {
	c := &Console{
		useColor: true,
	}
	for _, o := range opt {
		o(c)
	}

	return c
}

func (c *Console) DisableColor() {
	c.useColor = false
}

func (c *Console) DisableBorder() {
	c.opt = append(c.opt, text.DisableBorder())
}

func (c *Console) Info(format string, v ...interface{}) {
	msg := fmt.Sprintf(format, v...)
	opt := []text.Option{text.WithPrefix("[INFO] ", time.Now().Format(dateFormat)+" ")}
	opt = append(opt, c.opt...)
	output := text.Convert(msg, opt...)
	if runtime.GOOS != "windows" && c.useColor {
		output = aurora.Green(output).String()
	}
	c.Fprintf(output)
}

func (c *Console) Debug(format string, v ...interface{}) {
	msg := fmt.Sprintf(format, v...)
	opt := []text.Option{text.WithPrefix("[DEBUG] ", time.Now().Format(dateFormat)+" ")}
	opt = append(opt, c.opt...)
	output := text.Convert(msg, opt...)
	if runtime.GOOS != "windows" && c.useColor {
		output = aurora.Blue(output).String()
	}
	c.Fprintf(output)
}

func (c *Console) Warn(format string, v ...interface{}) {
	msg := fmt.Sprintf(format, v...)
	opt := []text.Option{text.WithPrefix("[WARN] ", time.Now().Format(dateFormat)+" ")}
	opt = append(opt, c.opt...)
	output := text.Convert(msg, opt...)
	if runtime.GOOS != "windows" && c.useColor {
		output = aurora.Yellow(output).String()
	}
	c.Fprintf(output)
}

func (c *Console) Error(err error) {
	err = errors.WithStack(err)
	msg := fmt.Sprintf("%+v", err)
	opt := []text.Option{text.WithPrefix("[ERROR] ", time.Now().Format(dateFormat)+" ")}
	opt = append(opt, c.opt...)
	output := text.Convert(msg, opt...)
	if runtime.GOOS != "windows" && c.useColor {
		output = aurora.Red(output).String()
	}
	c.Fprintf(output, true)
}

func (c *Console) Fprintf(msg string, err ...bool) {
	if len(c.output) == 0 {
		if len(err) > 0 && err[0] {
			fmt.Fprint(os.Stderr, msg)
		} else {
			fmt.Fprint(os.Stdout, msg)
		}
		return
	}

	c.saveFile(msg, err...)
}

func (c *Console) saveFile(msg string, error ...bool) {
	// TODO
}

func (c *Console) Rotate() error {
	panic("implement me")
}
