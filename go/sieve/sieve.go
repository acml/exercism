// Package sieve provides a function find all the primes from 2 up to a given
// number.
package sieve

import "math"

// Sieve of Eratosthenes is a simple, ancient algorithm for finding all prime
// numbers up to any given limit. It does so by iteratively marking as composite
// (i.e. not prime) the multiples of each prime, starting with the multiples of
// 2. It does not use any division or remainder operation.
func Sieve(limit int) []int {
	A := make([]bool, limit+1)
	for i := 2; i <= int(math.Sqrt(float64(limit))); i++ {
		if !A[i] {
			for j := i * i; j <= limit; j += i {
				A[j] = true
			}
		}
	}

	result := []int{}
	for i := 2; i <= limit; i++ {
		if !A[i] {
			result = append(result, i)
		}
	}
	return result
}
