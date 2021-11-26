package text

import (
	"strings"
	"unicode/utf8"
)

type Row struct {
	lineText  []string
	maxLength int
}

func NewRows(text ...string) []*Row {
	var list []*Row
	for _, t := range text {
		texts := strings.Split(t, "\n")
		var max = 0
		for _, t := range texts {
			runeCount := utf8.RuneCountInString(t)
			if max < runeCount {
				max = runeCount
			}
		}
		list = append(list, &Row{
			lineText:  texts,
			maxLength: max,
		})
	}
	return list
}
