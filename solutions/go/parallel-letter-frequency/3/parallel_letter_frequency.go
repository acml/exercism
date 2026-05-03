package letter

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

// ConcurrentFrequency makes concurrent calls to Frequency and combines results.
func ConcurrentFrequency(input []string) FreqMap {

	c := make(chan FreqMap, 10)
	for _, s := range input {
		go func(str string) {
			c <- Frequency(str)
		}(s)
	}

	res := FreqMap{}
	for range input {
		for k, v := range <-c {
			res[k] += v
		}
	}
	return res
}
