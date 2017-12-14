/*
Package luhn implements the Luhn formula
which can be used for validation purposes.
*/
package luhn

import (
	log "github.com/sirupsen/logrus"
	"strconv"
	"strings"
	"unicode"
)

// Valid returns true if the given string is a valid per the Luhn formula
func Valid(s string) bool {
	s = strings.Replace(s, " ", "", -1)
	if len(s) <= 1 {
		return false
	}

	sum := 0
	for i, c := range reverse(s) {
		if !unicode.IsDigit(c) {
			return false
		}

		x, err := strconv.Atoi(string(c))
		switch {
		case err != nil:
			log.WithError(err).Error("failed to convert string to int")
		case (i+1)%2 == 0:
			sum += x * 2
			if x*2 > 9 {
				sum -= 9
			}
		default:
			sum += x
		}
	}

	if sum%10 != 0 {
		return false
	}

	return true
}

func reverse(s string) string {
	var reverse string
	for i := len(s) - 1; i >= 0; i-- {
		reverse += string(s[i])
	}
	return reverse
}
