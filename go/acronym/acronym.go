/*
Package acronym implements logic for abbreviating long name into an acronym.
*/
package acronym

import (
	"strings"
)

// Abbreviate used to convert long name to an acronym
func Abbreviate(s string) string {
	var a string
	for _, l := range strings.FieldsFunc(s, Split) {
		a += l[0:1]
	}

	return strings.ToUpper(a)
}

// Split returns true when the given rune is either a whitespace or
// hyphen, indicating that should split on the given rune
func Split(r rune) bool {
	return r == ' ' || r == '-'
}
