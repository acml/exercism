package prime

// Factors computes the prime factors of a given natural number.
func Factors(input int64) []int64 {
	res := []int64{}
	for input > 1 {
		for i := int64(2); i <= input; i++ {
			if input%i == 0 {
				res = append(res, i)
				input /= i
				break
			}
		}
	}
	return res
}
