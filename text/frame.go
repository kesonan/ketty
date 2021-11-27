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
	"bytes"
	"fmt"
	"strings"
	"unicode/utf8"
)

// Frame is a text render with options, such as render with
// border, or with prefix.
type Frame struct {
	w        *bytes.Buffer
	rows     []*Row
	rowCount int
	maxWidth int
	border   BorderStyle
	prefix   []string
}

// Convert converts a plain text into another text with options.
func Convert(text string, opt ...Option) string {
	f := NewFrame(NewRows(text), opt...)
	f.draw()
	return f.w.String()
}

// NewFrame creates an instance of Frame.
func NewFrame(rows []*Row, opt ...Option) *Frame {
	f := &Frame{}
	options := []Option{WithLineStyle()}
	options = append(options, opt...)
	for _, o := range options {
		o(f)
	}

	var max int
	for _, r := range rows {
		if max < r.maxLength {
			max = r.maxLength
		}
	}

	f.w = bytes.NewBuffer(nil)
	f.rows = rows
	f.rowCount = len(rows)
	f.maxWidth = max
	return f
}

func (f *Frame) draw() {
	if len(f.rows) == 0 {
		return
	}
	f.drawTopBorder()
	f.drawRows()
	f.drawBottomBorder()
}

func (f *Frame) drawTopBorder() {
	switch f.border.(type) {
	case *NoneBorder:
		return
	}
	f.drawPrefix()
	f.w.WriteString(f.border.Corner().leftTop)
	f.drawSplitter()
	f.newLine()
}

func (f *Frame) drawPrefix() {
	for _, p := range f.prefix {
		f.w.WriteString(p)
	}
}

func (f *Frame) drawRows() {
	for index, r := range f.rows {
		f.drawText(r.lineText)
		if index < f.rowCount-1 {
			f.drawRowSplitter()
		}

	}
}

func (f *Frame) drawSplitter() {
	f.w.WriteString(f.drawHorizontalLine(f.maxWidth))
}

func (f *Frame) drawRowSplitter() {
	f.drawPrefix()
	f.w.WriteString(f.border.LineSplitter().verticalLine)
	f.w.WriteString(f.drawHorizontalLine(f.maxWidth))
	f.newLine()
}

func (f *Frame) drawText(texts []string) {
	for _, item := range texts {
		f.drawPrefix()
		f.w.WriteString(f.border.LineSplitter().verticalLine)
		switch f.border.(type) {
		case *NoneBorder:
			if len(f.prefix) > 0 {
				f.w.WriteString(" " + item)
			} else {
				f.w.WriteString(item)
			}
		default:
			f.w.WriteString("  " + item)

		}

		f.drawSpace(f.maxWidth - len(item))
		f.newLine()
	}
}

func (f *Frame) drawSpace(count int) {
	for i := 0; i < count; i++ {
		f.w.WriteString(" ")
	}
}

func (f *Frame) drawBottomBorder() {
	switch f.border.(type) {
	case *NoneBorder:
		return
	}
	f.drawPrefix()
	f.w.WriteString(f.border.Corner().leftBottom)
	f.drawSplitter()
	f.newLine()
}

func (f *Frame) drawHorizontalLine(count int) string {
	if len(f.border.LineSplitter().horizontalLine) == 0 {
		return ""
	}

	var line []string
	for {
		if len(line)*utf8.RuneCountInString(f.border.LineSplitter().horizontalLine) > count {
			break
		}
		line = append(line, f.border.LineSplitter().horizontalLine)
	}
	return strings.Join(line, "")
}

func (f *Frame) newLine() {
	f.w.WriteRune('\n')
}

// Print prints a text on terminal.
func (f *Frame) Print() {
	f.draw()
	fmt.Println(f.w.String())
}
