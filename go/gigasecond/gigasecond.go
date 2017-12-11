/*
Package gigasecond package implements gigasecond addition.
*/
package gigasecond

import (
	"math"
	"time"
)

// AddGigasecond adds a gigasecond to given time
func AddGigasecond(t time.Time) time.Time {
	return t.Add(time.Duration(math.Pow10(9)) * time.Second)
}
