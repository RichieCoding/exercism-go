// Package bob utlizies the Hey function to return string responses
package bob

import (
	"regexp"
	"strings"
)

const (
	questionResponse = "Sure."
	yellResponse = "Whoa, chill out!"
	yellQuestionResponse = "Calm down, I know what I'm doing!"
	emptyResponse = "Fine. Be that way!"
	defaultResponse = "Whatever."
)

// Hey returns responses based on the argument it recieves.
func Hey(remark string) string {
	remark = strings.TrimSpace(remark)
	question := strings.HasSuffix(remark, "?")
	yelling, _ := regexp.MatchString("\\A[A-Z1-9\\s[:punct:]]+[A-Z][A-Z1-9\\s[:punct:]]+\\z", remark)

	switch {
	case question && yelling:
		return yellQuestionResponse
	case !question && yelling:
		return yellResponse
	case question && !yelling:
		return questionResponse
	case remark == "":
		return emptyResponse
	default:
		return defaultResponse
	}

}
