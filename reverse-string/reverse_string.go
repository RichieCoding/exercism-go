package reverse

import "bytes"

// Reverse takes a word and reverses the order of each letter
func Reverse(word string) string {
	str := []rune(word)
	var buf bytes.Buffer

	for i := len(str); i > 0; i-- {
		buf.WriteRune(str[i-1])
	}

	return buf.String()
}
