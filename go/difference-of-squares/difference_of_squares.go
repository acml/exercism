// Package diffsquares provides a functions to find the difference between the
// square of the sum and the sum of the squares of the first N natural numbers.
package diffsquares

// SquareOfSum calculates the square of the sum of the first ten natural
// numbers.
// https://www.youtube.com/watch?v=ZplzrdKarX4
func SquareOfSum(n int) int {
	sum := n * (n + 1) / 2
	return sum * sum
}

// SumOfSquares calculates the sum of the squares of the first ten natural
// numbers.
// References to understand the formula used:
// https://www.youtube.com/watch?v=i7iKLZQ-vCk
// https://www.youtube.com/watch?v=MkGXR8umLco
// https://www.youtube.com/watch?v=KvvFTPsZxU4&t=191s
func SumOfSquares(n int) int {
	return n * (n + 1) * (2*n + 1) / 6
}

// Difference finds the difference between the square of the sum and the sum of
// the squares of the first N natural numbers.
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
