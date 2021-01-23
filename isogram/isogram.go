package isogram

import (
	"unicode"
)

// IsIsogram determines if a string contains any repeating letters, if so it is considered an isogram
func IsIsogram(s string) bool {
	var occurances = make(map[rune]bool)

	for _, letter := range s {
		letter = unicode.ToUpper(letter)
		if letter == '-' || letter == ' ' {
			continue
		}
		if occurances[letter] {
			return false
		}
		occurances[letter] = true
	}

	return true
}
