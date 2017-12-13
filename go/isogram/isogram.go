/*
Package isogram implements logic for determining
if the phrase is an isogram.
*/
package isogram

import (
	"strings"
	"unicode"
)

// IsIsogram returns true if the given word is an isogram
func IsIsogram(w string) bool {
	letters := make(map[string]bool)
	for _, l := range strings.ToLower(w) {
		letter := string(l)
		switch {
		case !unicode.IsLetter(l):
			continue
		case letters[letter]:
			return false
		}
		letters[letter] = true
	}
	return true
}
