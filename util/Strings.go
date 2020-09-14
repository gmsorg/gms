package util

import (
	"unicode"
	"unicode/utf8"
)

func IsExported(name string) bool {
	rune, _ := utf8.DecodeRuneInString(name)
	return unicode.IsUpper(rune)
}
