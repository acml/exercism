package piglatin

import "strings"

// Sentence converts input strings to Pig Latin which is a made-up children's
// language that's intended to be confusing.
func Sentence(input string) string {
	inputs := strings.Fields(input)

	res := []string{}
	for _, input := range inputs {
		if strings.HasPrefix(input, "xr") || strings.HasPrefix(input, "yt") || strings.IndexAny(input, "aeiou") == 0 {
			res = append(res, input+"ay")
			continue
		}

		index := strings.IndexFunc(input, func(r rune) bool { return strings.ContainsRune("aeio", r) })
		if index < 0 {
			index = strings.IndexFunc(input, func(r rune) bool { return strings.ContainsRune("uy", r) })
		}
		res = append(res, input[index:]+input[:index]+"ay")
	}
	return strings.Join(res, " ")
}
