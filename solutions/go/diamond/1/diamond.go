package diamond

import (
	"errors"
	"fmt"
	"strings"
)

// Gen takes as its input a letter, and outputs it in a diamond shape.
func Gen(b byte) (string, error) {
	if b < 'A' || b > 'Z' {
		return "", errors.New("Invalid Char")
	}

	l := int(b-'A') + 1
	sb := strings.Builder{}
	for i := 0; i < l-1; i++ {
		for j := 0; j < 2*l-1; j++ {
			if j == (2*l-1)/2-i || j == (2*l-1)/2+i {
				fmt.Fprintf(&sb, "%c", 'A'+(i%l))
			} else {
				fmt.Fprintf(&sb, " ")
			}
		}
		fmt.Fprintf(&sb, "\n")
	}
	for i := l - 1; i >= 0; i-- {
		for j := 0; j < 2*l-1; j++ {
			if j == (2*l-1)/2-i || j == (2*l-1)/2+i {
				fmt.Fprintf(&sb, "%c", 'A'+(i%l))
			} else {
				fmt.Fprintf(&sb, " ")
			}
		}
		fmt.Fprintf(&sb, "\n")
	}
	return sb.String(), nil
}
