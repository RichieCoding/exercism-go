// Package twofer utilizes ShareWith function
package twofer

import "fmt"

// ShareWith returns "One for <name>, and one for me."
// If an empty string is provided it replaces <name> with "you"
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}
	return fmt.Sprintf("One for %s, one for me.", name)
}
