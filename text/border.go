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

// Corner describes the character for left side corner.
type Corner struct {
	leftTop    string
	leftBottom string
}

// LineSplitter describes the splitter in horizontal and vertica direction.
type LineSplitter struct {
	horizontalLine string
	verticalLine   string
}

// BorderStyle is an interface which describes a border style
// for log print.
type BorderStyle interface {
	Corner() Corner
	LineSplitter() LineSplitter
}

// Border implements BorderStyle to receive a border style.
type Border struct {
	c Corner
	l LineSplitter
}

// Corner is the corner of border.
func (b *Border) Corner() Corner {
	return b.c
}

// LineSplitter describes a line splitter in horizontal and vertical.
func (b *Border) LineSplitter() LineSplitter {
	return b.l
}

// NoneBorder implements BorderStyle represents borderless.
type NoneBorder struct{}

// Corner is the corner of border.
func (b *NoneBorder) Corner() Corner {
	return Corner{}
}

// LineSplitter describes a line splitter in horizontal and vertical.
func (b *NoneBorder) LineSplitter() LineSplitter {
	return LineSplitter{}
}

// WithLineStyle describes a line style for border.
func WithLineStyle() Option {
	return WithBorder("┌", "└", "─", "│")
}

// WithDotStyle describes a dot style for border.
func WithDotStyle() Option {
	return WithCommonBorder(".")
}

// WithStarStyle describes a star style for border.
func WithStarStyle() Option {
	return WithCommonBorder("*")
}

// WithPlusStyle describes a plus style for border.
func WithPlusStyle() Option {
	return WithBorder("+", "+", "-", "|")
}

// DisableBorder represents that is disable border.
func DisableBorder() Option {
	return func(f *Frame) {
		f.border = &NoneBorder{}
	}
}

// WithFivePointedStarStyle describes a file-pointed star style for border.
func WithFivePointedStarStyle() Option {
	return WithCommonBorder("★")
}

// WithDoubleLine describes a double line style for border.
func WithDoubleLine() Option {
	return WithBorder("╔", "╚", "═", "║")
}
