/*
Package clock implements clock that handles times
without dates.
*/
package clock

import (
	"fmt"
)

// CustomClock holds hours and minutes for the no
type CustomClock struct {
	Hours   int
	Minutes int
}

// New used to create a new instance of the custom clock
func New(hours int, minutes int) CustomClock {
	c := CustomClock{}
	c.Hours, c.Minutes = getTime(hours, minutes)
	return c
}

// Add used to add minutes to the custom clock
func (c CustomClock) Add(minutes int) CustomClock {
	c.Hours, c.Minutes = getTime(c.Hours, c.Minutes+minutes)
	return c
}

// String returns formatted custom clock time
func (c *CustomClock) String() string {
	return fmt.Sprintf("%02d:%02d", c.Hours, c.Minutes)
}

func getTime(h, m int) (int, int) {
	minutes := m % 60
	if minutes < 0 {
		minutes += 60
		h--
	}

	hours := (h + (m / 60)) % 24
	if hours < 0 {
		hours += 24
	}

	return hours, minutes
}
