/*
Package raindrops provides a function to convert a number into a string that
contains raindrop sounds corresponding to certain potential factors.
*/
package raindrops

import "fmt"

/*
Convert returns a string that contains raindrop sounds corresponding to certain
potential factors.
*/
func Convert(input int) string {
	result := ""

	if input%3 == 0 {
		result = "Pling"
	}

	if input%5 == 0 {
		result += "Plang"
	}

	if input%7 == 0 {
		result += "Plong"
	}

	if len(result) == 0 {
		result = fmt.Sprintf("%d", input)
	}

	return result
}
