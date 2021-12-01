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
	"bufio"
	"fmt"
	"os"
	"runtime"
	"strings"
	"sync"
	"time"

	"github.com/anqiansong/ketty"
	"github.com/anqiansong/ketty/text"
	"github.com/logrusorgru/aurora"
	"github.com/pkg/errors"
)

const dateFormat = `2006-01-02 15:04:05.999`

// Option is an alias of the function with argument Console.
type Option func(c *Console)

var _ ketty.Logger = (*Console)(nil)

// Console is log printer which implements Logger.
type Console struct {
	useColor  bool
	usePrefix bool
	opt       []text.Option
	output    string
	testTime  string
	prefix    string
	w         *bufio.Writer
	ticker    *time.Ticker
	lock      sync.Mutex
	once      sync.Once
	doneChan  chan struct{}
}

// NewConsole creates an instance of Console.
func NewConsole(opt ...Option) *Console {
	c := &Console{
		useColor:  true,
		usePrefix: true,
	}
	for _, o := range opt {
		o(c)
	}
	return c
}

// DisableColor sets the color flag as false, if it is,
// the Console won't print log with color.
func (c *Console) DisableColor() {
	c.useColor = false
}

// DisablePrefix prints log without prefix
func (c *Console) DisablePrefix() {
	c.usePrefix = false
}

// DisableBorder prints log with borderless.
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

// Info prints info level log.
func (c *Console) Info(format string, v ...interface{}) {
	msg := fmt.Sprintf(format, v...)
	var opt []text.Option
	if c.usePrefix {
		if len(c.prefix) > 0 {
			opt = []text.Option{text.WithPrefix(c.prefix)}
		} else {
			opt = []text.Option{text.WithPrefix("[INFO] ", c.now())}
		}
	}
	opt = append(opt, c.opt...)
	output := text.Convert(msg, opt...)
	if runtime.GOOS != "windows" && c.useColor {
		output = aurora.Green(output).String()
	}
	c.fPrintf(output)
}

// Debug prints debug level log.
func (c *Console) Debug(format string, v ...interface{}) {
	env := os.Getenv(debugKey)
	if strings.EqualFold(env, "false") {
		return
	}

	msg := fmt.Sprintf(format, v...)
	var opt []text.Option
	if c.usePrefix {
		if len(c.prefix) > 0 {
			opt = []text.Option{text.WithPrefix(c.prefix)}
		} else {
			opt = []text.Option{text.WithPrefix("[DEBUG] ", c.now())}
		}
	}
	opt = append(opt, c.opt...)
	output := text.Convert(msg, opt...)
	if runtime.GOOS != "windows" && c.useColor {
		output = aurora.Blue(output).String()
	}
	c.fPrintf(output)
}

// Warn prints warn level log.
func (c *Console) Warn(format string, v ...interface{}) {
	msg := fmt.Sprintf(format, v...)
	var opt []text.Option
	if c.usePrefix {
		if len(c.prefix) > 0 {
			opt = []text.Option{text.WithPrefix(c.prefix)}
		} else {
			opt = []text.Option{text.WithPrefix("[WARN] ", c.now())}
		}
	}
	opt = append(opt, c.opt...)
	output := text.Convert(msg, opt...)
	if runtime.GOOS != "windows" && c.useColor {
		output = aurora.Yellow(output).String()
	}
	c.fPrintf(output)
}

// Error prints error level log.
func (c *Console) Error(err error) {
	err = errors.WithStack(err)
	msg := fmt.Sprintf("%+v", err)
	var opt []text.Option
	if c.usePrefix {
		if len(c.prefix) > 0 {
			opt = []text.Option{text.WithPrefix(c.prefix)}
		} else {
			opt = []text.Option{text.WithPrefix("[ERROR] ", c.now())}
		}
	}
	opt = append(opt, c.opt...)
	output := text.Convert(msg, opt...)
	if runtime.GOOS != "windows" && c.useColor {
		output = aurora.Red(output).String()
	}
	c.fPrintf(output, true)
}

// ErrorText prints error level log.
func (c *Console) ErrorText(format string, v ...interface{}) {
	msg := fmt.Sprintf(format, v...)
	var opt []text.Option
	if c.usePrefix {
		if len(c.prefix) > 0 {
			opt = []text.Option{text.WithPrefix(c.prefix)}
		} else {
			opt = []text.Option{text.WithPrefix("[ERROR] ", c.now())}
		}
	}
	opt = append(opt, c.opt...)
	output := text.Convert(msg, opt...)
	if runtime.GOOS != "windows" && c.useColor {
		output = aurora.Red(output).String()
	}
	c.fPrintf(output)
}

func (c *Console) fPrintf(msg string, err ...bool) {
	if c.w == nil {
		if len(err) > 0 && err[0] {
			_, _ = fmt.Fprint(os.Stderr, msg)
		} else {
			_, _ = fmt.Fprint(os.Stdout, msg)
		}
		return
	}
	_, _ = fmt.Fprint(c.w, msg)
}

func (c *Console) Flush() error {
	if c.w == nil {
		return nil
	}
	c.lock.Lock()
	defer c.lock.Unlock()
	fmt.Println(time.Now())
	fmt.Println(c.w.Buffered())
	return c.w.Flush()
}

// rotate makes a new log file in interval.
func (c *Console) rotate() error {
	panic("implement me")
}

func (c *Console) Close() {
	c.once.Do(func() {
		_ = c.Flush()
		if c.ticker != nil {
			c.ticker.Stop()
		}
		if c.doneChan != nil {
			close(c.doneChan)
		}
	})
}

func (c *Console) listenFile() {
	for {
		select {
		case <-c.ticker.C:
			_ = c.Flush()
		case <-c.doneChan:
			return
		default:

		}
	}
}
