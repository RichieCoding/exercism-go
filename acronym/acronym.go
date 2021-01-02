// Package acronym creates acronyms from a string
package acronym

import (
	"strings"
)

// Abbreviate returns the first letter of each word in a phrase
func Abbreviate(s string) string {
	var acronym string
	s = strings.ReplaceAll(s, "-", " ")
	s = strings.ReplaceAll(s, "_", " ")
	phrase := strings.Fields(s)

	for _, word := range phrase {
		acronym += strings.ToUpper(word[0:1])
	}

	return acronym
}
