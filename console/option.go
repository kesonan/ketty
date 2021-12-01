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
	"os"
	"path/filepath"
	"time"

	"github.com/anqiansong/ketty/text"
)

// WithTextOption returns an Option wrapped text.Option.
func WithTextOption(opt ...text.Option) Option {
	return func(c *Console) {
		c.opt = opt
	}
}

// WithOutputDir sets an output file for Logger.
func WithOutputDir(file string) Option {
	return func(c *Console) {
		c.lock.Lock()
		defer c.lock.Unlock()
		if len(file) == 0 {
			return
		}
		info, err := os.Stat(file)
		if err != nil && os.IsNotExist(err) {
			err = os.MkdirAll(file, os.ModePerm)
			if err != nil {
				panic(err)
			}
		} else {
			if !info.IsDir() {
				panic("invalid dir: " + file)
			}
		}

		c.output = file
		filename := filepath.Join(c.output, "access.log")
		f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
		if err != nil {
			panic(err)
		}
		c.w = bufio.NewWriter(f)
		c.ticker = time.NewTicker(time.Second)
		c.doneChan = make(chan struct{})
		go c.listenFile()
	}
}

// WithPrefix sets a prefix for Logger.
func WithPrefix(prefix string) Option {
	return func(c *Console) {
		c.prefix = prefix
	}
}
