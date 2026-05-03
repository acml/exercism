package diamond

import (
	"errors"
	"fmt"
	"io"
	"strings"
)

// Gen takes as its input a letter, and outputs it in a diamond shape.
func Gen(b byte) (string, error) {
	if b < 'A' || b > 'Z' {
		return "", errors.New("Invalid Char")
	}

	length := int(b-'A') + 1
	sb := strings.Builder{}
	for i := 0; i < length-1; i++ {
		diamondLine(&sb, i, length)
		fmt.Fprintf(&sb, "\n")
	}
	for i := length - 1; i >= 0; i-- {
		diamondLine(&sb, i, length)
		fmt.Fprintf(&sb, "\n")
	}
	return sb.String(), nil
}

func diamondLine(w io.Writer, n, length int) {
	for j := 0; j < 2*length-1; j++ {
		if j == (2*length-1)/2-n || j == (2*length-1)/2+n {
			fmt.Fprintf(w, "%c", 'A'+(n%length))
		} else {
			fmt.Fprintf(w, " ")
		}
	}
}
