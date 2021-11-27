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

package text

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBorder_Corner(t *testing.T) {
	b := Border{}
	assert.Equal(t, Corner{}, b.Corner())
}

func TestBorder_LineSplitter(t *testing.T) {
	b := Border{}
	assert.Equal(t, LineSplitter{}, b.LineSplitter())
}

func TestNoneBorder_Corner(t *testing.T) {
	b := NoneBorder{}
	assert.Equal(t, Corner{}, b.Corner())
}

func TestNoneBorder_LineSplitter(t *testing.T) {
	b := NoneBorder{}
	assert.Equal(t, LineSplitter{}, b.LineSplitter())
}

func TestWithLineStyle(t *testing.T) {
	o := WithLineStyle()
	f := NewFrame(NewRows("hello ketty"), o)
	expected := Border{
		c: Corner{
			leftTop:    "┌",
			leftBottom: "└",
		},
		l: LineSplitter{
			horizontalLine: "─",
			verticalLine:   "│",
		},
	}
	actual, ok := f.border.(*Border)
	assert.True(t, ok)
	assert.Equal(t, expected, *actual)
}

func TestWithDotStyle(t *testing.T) {
	o := WithDotStyle()
	f := NewFrame(NewRows("hello ketty"), o)
	expected := Border{
		c: Corner{
			leftTop:    ".",
			leftBottom: ".",
		},
		l: LineSplitter{
			horizontalLine: ".",
			verticalLine:   ".",
		},
	}
	actual, ok := f.border.(*Border)
	assert.True(t, ok)
	assert.Equal(t, expected, *actual)
}

func TestWithStarStyle(t *testing.T) {
	o := WithStarStyle()
	f := NewFrame(NewRows("hello ketty"), o)
	expected := Border{
		c: Corner{
			leftTop:    "*",
			leftBottom: "*",
		},
		l: LineSplitter{
			horizontalLine: "*",
			verticalLine:   "*",
		},
	}
	actual, ok := f.border.(*Border)
	assert.True(t, ok)
	assert.Equal(t, expected, *actual)
}

func TestWithPlusStyle(t *testing.T) {
	o := WithPlusStyle()
	f := NewFrame(NewRows("hello ketty"), o)
	expected := Border{
		c: Corner{
			leftTop:    "+",
			leftBottom: "+",
		},
		l: LineSplitter{
			horizontalLine: "-",
			verticalLine:   "|",
		},
	}
	actual, ok := f.border.(*Border)
	assert.True(t, ok)
	assert.Equal(t, expected, *actual)
}

func TestDisableBorder(t *testing.T) {
	o := DisableBorder()
	f := NewFrame(NewRows("hello ketty"), o)
	expected := NoneBorder{}
	actual, ok := f.border.(*NoneBorder)
	assert.True(t, ok)
	assert.Equal(t, expected, *actual)
}

func TestWithFivePointedStarStyle(t *testing.T) {
	o := WithFivePointedStarStyle()
	f := NewFrame(NewRows("hello ketty"), o)
	expected := Border{
		c: Corner{
			leftTop:    "★",
			leftBottom: "★",
		},
		l: LineSplitter{
			horizontalLine: "★",
			verticalLine:   "★",
		},
	}
	actual, ok := f.border.(*Border)
	assert.True(t, ok)
	assert.Equal(t, expected, *actual)
}

func TestWithDoubleLine(t *testing.T) {
	o := WithDoubleLine()
	f := NewFrame(NewRows("hello ketty"), o)
	expected := Border{
		c: Corner{
			leftTop:    "╔",
			leftBottom: "╚",
		},
		l: LineSplitter{
			horizontalLine: "═",
			verticalLine:   "║",
		},
	}
	actual, ok := f.border.(*Border)
	assert.True(t, ok)
	assert.Equal(t, expected, *actual)
}