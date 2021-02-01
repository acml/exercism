// Package grains provides functions to calculate the number of grains of wheat
// on a chessboard.
package grains

import "errors"

// Square calculates how many grains were on a given square.
func Square(n int) (uint64, error) {
	if n < 1 || n > 64 {
		return 0, errors.New("n must be a positive integer between 1 and 64")
	}
	return 1 << (n - 1), nil
}

// Total gives the total number of grains on the chessboard.
func Total() uint64 {
	var sum uint64
	for n := 1; n <= 64; n++ {
		sum += (1 << (n - 1))
	}
	return sum
}
