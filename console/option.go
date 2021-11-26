package console

import (
	"github.com/anqiansong/ketty/text"
)

func WithTextOption(opt ...text.Option) Option {
	return func(c *Console) {
		c.opt = opt
	}
}

func WithOutputDir(file string) Option {
	return func(c *Console) {
		c.output = file
	}
}
