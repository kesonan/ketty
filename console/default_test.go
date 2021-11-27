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

package console

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDisableColor(t *testing.T) {
	DisableColor()
	assert.False(t, colorConsole.useColor)
}

func ExampleDisableBorder() {
	DisableBorder()
	DisableColor()
	colorConsole.useTestTime("2021-11-27")
	Info("hello ketty")
	// Output:
	// [INFO] 2021-11-27 hello ketty
}

func ExampleInfo() {
	DisableColor()
	colorConsole.useTestTime("2021-11-27")
	Info("hello ketty")
	// Output:
	// [INFO] 2021-11-27 hello ketty
}

func ExampleDebug() {
	DisableColor()
	colorConsole.useTestTime("2021-11-27")
	Debug("hello ketty")
	// Output:
	// [DEBUG] 2021-11-27 hello ketty
}

func ExampleWarn() {
	DisableColor()
	colorConsole.useTestTime("2021-11-27")
	Warn("hello ketty")
	// Output:
	// [WARN] 2021-11-27 hello ketty
}

func Test_Error(t *testing.T) {
	DisableColor()
	colorConsole.useTestTime("2021-11-27")
	Error(errors.New("hello ketty"))
}
