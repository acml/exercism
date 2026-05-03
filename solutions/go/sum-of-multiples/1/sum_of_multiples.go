// Package summultiples provides a function that finds the sum of all the unique
// multiples of particular numbers.
package summultiples

// SumMultiples given a number, finds the sum of all the unique multiples of
// particular numbers up to but not including that number.
func SumMultiples(limit int, divisors ...int) int {
	sum := 0
	multiples := map[int]bool{}
	for _, divisor := range divisors {
		if divisor == 0 {
			continue
		}
		for i := 1; i*divisor < limit; i++ {
			multiples[divisor*i] = true
		}
	}

	for k, v := range multiples {
		if v {
			sum += k
		}
	}
	return sum
}
