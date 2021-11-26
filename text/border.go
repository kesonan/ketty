package text

type Corner struct {
	leftTop    string
	leftBottom string
}

type LineSplitter struct {
	horizontalLine string
	verticalLine   string
}

type BorderStyle interface {
	Corner() Corner
	LineSplitter() LineSplitter
}

type Border struct {
	c Corner
	l LineSplitter
}

func (b *Border) Corner() Corner {
	return b.c
}

func (b *Border) LineSplitter() LineSplitter {
	return b.l
}

func WithLineStyle() Option {
	return WithBorder("┌", "└", "─", "│")
}

func WithDotStyle() Option {
	return WithCommonBorder(".")
}

func WithStarStyle() Option {
	return WithCommonBorder("*")
}

func WithPlusStyle() Option {
	return WithBorder("+", "+", "-", "|")
}

func DisableBorder() Option {
	return WithCommonBorder("")
}

func WithFivePointedStarStyle() Option {
	return WithCommonBorder("★")
}

func WithDoubleLine() Option {
	return WithBorder("╔", "╚", "═", "║")
}
