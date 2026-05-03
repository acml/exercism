// Package clock provides clock related New, String, Add and Subtract functions.
package clock

import "fmt"

// Clock represents time in hours and minutes.
type Clock struct {
	hour   int
	minute int
}

func absMod(n, m int) int {
	return ((n % m) + m) % m
}

func clockToMinutes(hour, minute int) int {
	minutes := absMod(hour, 24)*60 + minute
	minutes = absMod(minutes, 24*60)
	return minutes
}

// New produces clock time from given hour and minutes.
func New(hour, minute int) Clock {
	minutes := clockToMinutes(hour, minute)
	return Clock{minutes / 60, minutes % 60}
}

// String outputs the clock time in HH:MM format.
func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hour, c.minute)
}

// Add adds minutes to clock time.
func (c Clock) Add(minute int) Clock {
	return c.Subtract(-minute)
}

// Subtract subtracts minutes from clock time.
func (c Clock) Subtract(minute int) Clock {
	minutes := clockToMinutes(c.hour, c.minute) - minute
	minutes = absMod(minutes, 24*60)
	return Clock{minutes / 60, minutes % 60}
}
