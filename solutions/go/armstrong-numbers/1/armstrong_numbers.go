package armstrong

import "math"

// IsNumber checks if given input is a number that is the sum of its own digits
// each raised to the power of the number of digits.
func IsNumber(n int) bool {
	ndigits := 0
	digits := []int{}
	for x := n; x > 0; ndigits++ {
		digits = append(digits, x%10)
		x /= 10
	}

	sum := 0
	for _, digit := range digits {
		sum += int(math.Pow(float64(digit), float64(ndigits)))
	}
	return sum == n
}
