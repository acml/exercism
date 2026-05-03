package series

// All returns a list of all substrings of s with length n.
func All(n int, s string) []string {

	res := []string{}
	for i := 0; i+n <= len(s); i++ {
		res = append(res, s[i:i+n])
	}
	return res
}

// UnsafeFirst returns the first substring of s with length n.
func UnsafeFirst(n int, s string) string {
	return s[:n]
}

// First returns the first substring of s with length n.
func First(n int, s string) (first string, ok bool) {
	if n > len(s) {
		return "", false
	}
	return s[:n], true
}
