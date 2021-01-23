package isogram

import (
	"strings"
	"unicode"
)

// IsIsogram determines if a string contains any repeating letters, if so it is considered an isogram
func IsIsogram(s string) bool {
	s = strings.ReplaceAll(s, "-", "")
	s = strings.ReplaceAll(s, " ", "")

	var occurances = make(map[rune]bool)

	for _, letter := range s {
		letter = unicode.ToUpper(letter)
		if occurances[letter] {
			return false
		}
		occurances[letter] = true
	}

	return true
}
