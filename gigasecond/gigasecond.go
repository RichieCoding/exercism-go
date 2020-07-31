// Package gigasecond implemnets AddGigasecond function
package gigasecond

// import path for the time package from the standard library
import "time"

// AddGigasecond adds a gigasecond to the current time you have
func AddGigasecond(t time.Time) time.Time {
	return t.Add(time.Second * 1000000000)
}

