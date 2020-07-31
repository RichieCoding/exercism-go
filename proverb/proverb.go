// Package proverb takes a list of words and creates proverbs
package proverb

import "fmt"

const (
	text = "For want of a %s the %s was lost."
	last = "And all for the want of a %s."
)

// Proverb takes a list of words and returns a list of proverbs
func Proverb(rhyme []string) []string {
	var ft []string

	if len(rhyme) <= 0 {
		return ft
	}

	for i := 0; i < len(rhyme) - 1; i++ {
		ft = append(ft, fmt.Sprintf(text, rhyme[i], rhyme[i + 1]))
	}
	ft = append(ft, fmt.Sprintf(last, rhyme[0]))
	return ft
}

