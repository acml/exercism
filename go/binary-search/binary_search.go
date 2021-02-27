package binarysearch

// SearchInts implements a binary search algorithm.
func SearchInts(inputs []int, key int) int {
	l, r := 0, len(inputs)-1
	for l <= r {
		m := (l + r) / 2
		if inputs[m] < key {
			l = m + 1
			continue
		}

		if inputs[m] > key {
			r = m - 1
			continue
		}

		return m
	}

	return -1
}
