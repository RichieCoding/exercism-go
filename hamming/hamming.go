// Package hamming utilizes Distance function that returns an int and error
package hamming

import "errors"

// Distance function takes in two strings and compares the differences between one another with each letter
// Returns the amount of differences in an int and an error
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("The DNA's do not match")
	}
	var sum int
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			sum++
		}
	}
	return sum, nil
}
