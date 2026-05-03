package letter

import "sync"

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

// ConcurrentFrequency does the frequency count concurrently for each item of
// input slice of strings
func ConcurrentFrequency(input []string) FreqMap {
	var wg sync.WaitGroup
	wg.Add(len(input))

	maps := make([]map[rune]int, len(input))
	for i, s := range input {
		go func(m *map[rune]int, s string) {
			*m = Frequency(s)
			wg.Done()
		}(&maps[i], s)
	}
	wg.Wait()

	res := map[rune]int{}
	for _, m := range maps {
		for k, v := range m {
			res[k] += v
		}
	}
	return res
}
