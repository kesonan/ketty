package text

import (
	"bytes"
	"fmt"
	"strings"
	"unicode/utf8"
)

type Frame struct {
	w        *bytes.Buffer
	rows     []*Row
	rowCount int
	maxWidth int
	border   BorderStyle
	prefix   []string
}

func Convert(text string, opt ...Option) string {
	f := NewFrame(NewRows(text), opt...)
	f.draw()
	return f.w.String()
}

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
	f.newLine()
	f.drawRows()
	f.drawBottomBorder()
	f.newLine()
}

func (f *Frame) drawTopBorder() {
	f.drawPrefix()
	f.w.WriteString(f.border.Corner().leftTop)
	f.drawSplitter()
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
		f.w.WriteString("  " + item)
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
	f.drawPrefix()
	f.w.WriteString(f.border.Corner().leftBottom)
	f.drawSplitter()
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

func (f *Frame) Print() {
	f.draw()
	fmt.Println(f.w.String())
}
