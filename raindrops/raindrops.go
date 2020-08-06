// Package raindrops converts a number to a string depending on the numbers factors
package raindrops

import "strconv"

// Convert recieves an int and depending on the int will return Pling, Plang, Plong or the original int
func Convert(n int) string {
	var raindrop string

	if n%3 == 0 {
		raindrop += "Pling"
	}

	if n%5 == 0 {
		raindrop += "Plang"
	}

	if n%7 == 0 {
		raindrop += "Plong"
	}

	if raindrop == "" {
		raindrop = strconv.Itoa(n)
	}

	return raindrop
}
