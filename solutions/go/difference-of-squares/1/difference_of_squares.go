// Package diffsquares provides a functions to find the difference between the
// square of the sum and the sum of the squares of the first N natural numbers.
package diffsquares

// SquareOfSum calculates the square of the sum of the first ten natural
// numbers.
func SquareOfSum(n int) int {
	var sum int
	for i := 1; i <= n; i++ {
		sum += i
	}
	return sum * sum
}

// SumOfSquares calculates the sum of the squares of the first ten natural
// numbers.
func SumOfSquares(n int) int {
	var sum int
	for i := 1; i <= n; i++ {
		sum += i * i
	}
	return sum
}

// Difference finds the difference between the square of the sum and the sum of
// the squares of the first N natural numbers.
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
