package text

type Option func(f *Frame)

func WithBorder(leftTop, leftBottom, horizontal, vertical string) Option {
	border := &Border{
		c: Corner{
			leftTop:    leftTop,
			leftBottom: leftBottom,
		},
		l: LineSplitter{
			horizontalLine: horizontal,
			verticalLine:   vertical,
		},
	}
	return func(f *Frame) {
		f.border = border
	}
}

func WithCommonBorder(char string) Option {
	border := &Border{
		c: Corner{
			leftTop:    char,
			leftBottom: char,
		},
		l: LineSplitter{
			horizontalLine: char,
			verticalLine:   char,
		},
	}
	return func(f *Frame) {
		f.border = border
	}
}

func WithPrefix(prefix ...string) Option {
	return func(f *Frame) {
		f.prefix = prefix
	}
}
