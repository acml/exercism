// Package prime provides a function which determines what the nth prime is.
package prime

func isPrime(n int) bool {
	if n <= 3 {
		return n > 1
	}

	if n%2 == 0 || n%3 == 0 {
		return false
	}

	for i := 5; i*i <= n; i += 6 {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}
	return true
}

// Nth given a number n, determines what the nth prime is.
func Nth(n int) (int, bool) {
	count := 0
	for i := 2; i < 10000000; i++ {
		if isPrime(i) {
			count++
			if count == n {
				return i, true
			}
		}
	}
	return 0, false
}
