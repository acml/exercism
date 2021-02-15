package collatzconjecture

import "errors"

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
