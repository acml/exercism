package railfence

import "strings"

// Encode is a Rail Fence cipher encoder
func Encode(input string, rails int) string {
	var sb strings.Builder
	rollover := 2 * (rails - 1)
	for i := 0; i < rails; i++ {
		offset := rollover - 2*(i%(rails-1))
		for j := i; j < len(input); {
			sb.WriteByte(input[j])
			j = j + offset
			if offset != rollover {
				offset = rollover - offset
			}
		}
	}

	return sb.String()
}

// Decode is a Rail Fence cipher decoder
func Decode(input string, rails int) string {
	r := make([]rune, len(input))
	rollover := 2 * (rails - 1)
	pos := 0
	for i := 0; i < rails; i++ {
		offset := rollover - 2*(i%(rails-1))
		for j := i; j < len(input); {
			r[j] = rune(input[pos])
			pos++
			j = j + offset
			if offset != rollover {
				offset = rollover - offset
			}
		}
	}

	return string(r)
}
