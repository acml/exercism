// Package clock provides clock related New, String, Add and Subtract functions.
package clock

import "fmt"

// Clock represents time in minutes.
type Clock struct {
	minute int
}

// New produces clock time from given hour and minutes.
func New(hour, minute int) Clock {
	m := minute + hour*60
	m %= 24 * 60
	if m < 0 {
		m += 24 * 60
	}
	return Clock{m}
}

// String outputs the clock time in HH:MM format.
func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.minute/60, c.minute%60)
}

// Add adds minutes to clock time.
func (c Clock) Add(minute int) Clock {
	return New(0, c.minute+minute)
}

// Subtract subtracts minutes from clock time.
func (c Clock) Subtract(minute int) Clock {
	return New(0, c.minute-minute)
}
