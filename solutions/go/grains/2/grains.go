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
// Total progression:
// 1 (1 = 2 * (1 << (1 - 1)) - 1) -> 2 * Square(1) - 1
// 3 (3 = 2 * (1 << (2 - 1)) - 1) -> 2 * Square(2) - 1
// 7 (7 = 2 * (1 << (3 - 1)) - 1) -> 2 * Square(3) - 1
// ...
// x (x = 2 * (1 << (n - 1)) - 1)) -> 2 * Square(n) - 1
func Total() uint64 {
	return 2*(1<<(64-1)) - 1
}
