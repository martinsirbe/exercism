/*
Package raindrops implements logic for converting
integer to a string based on number's factors.
*/
package raindrops

import (
	"strconv"
)

// Raindrop contains factor and sound used for integer conversion
type Raindrop struct {
	Factor int
	Sound  string
}

var rs = []Raindrop{{3, "Pling"}, {5, "Plang"}, {7, "Plong"}}

// Convert converts given integer to a string depending on number's factors
func Convert(i int) string {
	var o string
	for _, r := range rs {
		if i%r.Factor == 0 {
			o += r.Sound
		}
	}

	switch {
	case o != "":
		return o
	default:
		return strconv.Itoa(i)
	}
}
