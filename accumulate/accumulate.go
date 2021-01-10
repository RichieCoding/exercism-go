package accumulate

import "strings"

// Accumulate recieves a collection of strings and returns each string uppercased
func Accumulate(s []string, upper func(string)string) []string {
	var ns []string

	if len(s) != 0 {
		for _, word := range s {
			ns = append(ns, upper(word))
		}
	}

	return ns
}

func toUppercase(s string) string {
	return strings.ToUpper(s)
}
