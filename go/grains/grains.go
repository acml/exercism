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
// 1 (1 = 2 ^ 1 - 1) ->
// 3 (1 + 2 = 2 ^ 2 - 1) ->
// 7 (1 + 2 + 4 = 2 ^ 3 - 1) ->
// ...
// x (1 + 2 + 4 + 8 + ... = 2 ^ 64 - 1)
func Total() uint64 {
	return 1<<64 - 1
}
