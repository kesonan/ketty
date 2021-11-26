package text

import (
	"testing"
)

func TestNewRow(t *testing.T) {
	rows := NewRows(`{
	"name": "k",
	"age": 20,
	"gender": "男"
}`)
	f := NewFrame(rows, WithPrefix("[INFO] ","2021-11-26 16:11 "), DisableBorder())
	f.Print()
}
