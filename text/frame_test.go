/*
 * MIT License
 *
 * Copyright (c) 2021 anqiansong
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */

package text

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewRow(t *testing.T) {
	rows := NewRows(`{
	"name": "k",
	"age": 20,
	"gender": "男"
}`)
	f := NewFrame(rows, WithPrefix("[INFO] ", "2021-11-26 16:11 "), DisableBorder())
	f.Print()
}

func TestConvert(t *testing.T) {
	actual := Convert("hello ketty", DisableBorder())
	assert.Equal(t, "hello ketty\n", actual)
}

func Test_draw(t *testing.T) {
	f:=NewFrame(NewRows("hello ketty"))
	t.Run("drawSplitter", func(t *testing.T) {
		f.drawSplitter()
	})

	t.Run("drawRowSplitter", func(t *testing.T) {
		f.drawRowSplitter()
	})

	t.Run("drawText", func(t *testing.T) {
		f.drawText([]string{})
	})

	t.Run("drawPrefix", func(t *testing.T) {
		f.drawPrefix()
	})

	t.Run("drawHorizontalLine", func(t *testing.T) {
		f.drawHorizontalLine(0)
	})
}
