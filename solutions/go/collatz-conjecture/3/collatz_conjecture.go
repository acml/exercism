package collatzconjecture

import "errors"

// CollatzConjecture performs The Collatz Conjecture
// Take any positive integer n. If n is even, divide n by 2 to get n / 2. If n is
// odd, multiply n by 3 and add 1 to get 3n + 1. Repeat the process indefinitely.
// The conjecture states that no matter which number you start with, you will
// always reach 1 eventually.
func CollatzConjecture(n int) (int, error) {

	if n <= 0 {
		return 0, errors.New("n must be positive integer")
	}

	var i int
	for ; n != 1; i++ {
		if n%2 == 0 {
			n /= 2
		} else {
			n = n*3 + 1
		}
	}

	return i, nil
}
