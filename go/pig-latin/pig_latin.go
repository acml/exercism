package piglatin

import "strings"

// Sentence converts input strings to Pig Latin which is a made-up children's
// language that's intended to be confusing.
func Sentence(input string) string {
	words := strings.Fields(input)

	res := []string{}
	for _, word := range words {
		if strings.HasPrefix(word, "xr") || strings.HasPrefix(word, "yt") || strings.IndexAny(word, "aeiou") == 0 {
			res = append(res, word+"ay")
			continue
		}

		index := strings.IndexFunc(word, func(r rune) bool { return strings.ContainsRune("aeio", r) })
		if index < 0 {
			index = strings.IndexFunc(word, func(r rune) bool { return strings.ContainsRune("uy", r) })
		}
		res = append(res, word[index:]+word[:index]+"ay")
	}
	return strings.Join(res, " ")
}
