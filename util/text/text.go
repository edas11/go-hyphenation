package text

import (
	"strings"
	"unicode"
)

func IndexFrom(str, search string, from int) int{
	if from > len(str) - 1 {
		return -1
	}
	matchInSubstring := strings.Index(str[from:], search)
	if matchInSubstring == -1 {
		return -1
	}
	return matchInSubstring + from
}

func RemoveDigit(character rune) rune {
	if unicode.IsDigit(character) || character == '.' {
		return rune(-1)
	}
	return character
}