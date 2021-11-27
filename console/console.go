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
	testTime string
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

func (c *Console) useTestTime(t string) {
	c.testTime = t
}

func (c *Console) now() string {
	now := time.Now().Format(dateFormat)
	if len(c.testTime) > 0 {
		now = c.testTime
	}
	return now
}

func (c *Console) Info(format string, v ...interface{}) {
	msg := fmt.Sprintf(format, v...)
	opt := []text.Option{text.WithPrefix("[INFO] ", c.now())}
	opt = append(opt, c.opt...)
	output := text.Convert(msg, opt...)
	if runtime.GOOS != "windows" && c.useColor {
		output = aurora.Green(output).String()
	}
	c.Fprintf(output)
}

func (c *Console) Debug(format string, v ...interface{}) {
	msg := fmt.Sprintf(format, v...)
	opt := []text.Option{text.WithPrefix("[DEBUG] ", c.now())}
	opt = append(opt, c.opt...)
	output := text.Convert(msg, opt...)
	if runtime.GOOS != "windows" && c.useColor {
		output = aurora.Blue(output).String()
	}
	c.Fprintf(output)
}

func (c *Console) Warn(format string, v ...interface{}) {
	msg := fmt.Sprintf(format, v...)
	opt := []text.Option{text.WithPrefix("[WARN] ", c.now())}
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
	opt := []text.Option{text.WithPrefix("[ERROR] ", c.now())}
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
