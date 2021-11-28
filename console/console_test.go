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
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewConsole(t *testing.T) {
	t.Run("withoutOption", func(t *testing.T) {
		c := NewConsole()
		assert.True(t, c.useColor)
		assert.True(t, len(c.output) == 0)
	})

	t.Run("withOption", func(t *testing.T) {
		c := NewConsole(WithOutputDir("foo"))
		assert.Equal(t, "foo", c.output)
	})
}

func TestConsole_DisableColor(t *testing.T) {
	c := NewConsole()
	c.DisableColor()
	assert.False(t, c.useColor)
}

func ExampleConsole_DisableBorder() {
	c := NewConsole()
	c.DisableBorder()
	c.DisableColor()
	c.useTestTime("2021-11-27")
	c.Info("hello ketty")
	// Output:
	// [INFO] 2021-11-27 hello ketty
}

func ExampleConsole_DisablePrefix() {
	c := NewConsole()
	c.DisableBorder()
	c.DisableColor()
	c.DisablePrefix()
	c.Info("hello ketty")
	// Output:
	// hello ketty
}

func ExampleConsole_Info() {
	c := NewConsole()
	c.DisableColor()
	c.useTestTime("2021-11-27")
	c.Info("hello ketty")
	// Output:
	// [INFO] 2021-11-27┌────────────
	// [INFO] 2021-11-27│  hello ketty
	// [INFO] 2021-11-27└────────────
}

func ExampleConsole_Debug() {
	c := NewConsole()
	c.DisableColor()
	c.useTestTime("2021-11-27")
	c.Debug("hello ketty")
	// Output:
	// [DEBUG] 2021-11-27┌────────────
	// [DEBUG] 2021-11-27│  hello ketty
	// [DEBUG] 2021-11-27└────────────
}

func ExampleConsole_Warn() {
	c := NewConsole()
	c.DisableColor()
	c.useTestTime("2021-11-27")
	c.Warn("hello ketty")
	// Output:
	// [WARN] 2021-11-27┌────────────
	// [WARN] 2021-11-27│  hello ketty
	// [WARN] 2021-11-27└────────────
}

func TestConsole_Error(t *testing.T) {
	c := NewConsole()
	c.DisableColor()
	c.useTestTime("2021-11-27")
	c.Error(errors.New("hello ketty"))
}

func ExampleConsole_ErrorText() {
	c := NewConsole()
	c.DisableColor()
	c.DisablePrefix()
	c.DisableBorder()
	c.ErrorText("hello ketty")
	// output:
	// hello ketty
}

func ExampleUsePrefix() {
	c := NewConsole()
	c.prefix = "test"
	c.DisableBorder()
	c.DisableColor()
	c.Info("hello ketty")
	// output:
	// test hello ketty
}
