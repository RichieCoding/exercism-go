package accumulate

// Accumulate recieves a collection of strings and returns each string uppercased
func Accumulate(s []string, upper func(string) string) []string {
	var ns []string

	for _, word := range s {
		ns = append(ns, upper(word))
	}

	return ns
}
