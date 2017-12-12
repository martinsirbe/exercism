/*
Package twofer implements two for one logic.
*/
package twofer

import "fmt"

// ShareWith returns a 2-fer sentence either using the given name
// or 'you' if no name is given
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}
	return fmt.Sprintf("One for %s, one for me.", name)
}
