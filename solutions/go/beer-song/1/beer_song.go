package beer

import (
	"errors"
	"fmt"
	"strings"
)

// Verse returns a specific verse of the beer song.
func Verse(n int) (string, error) {
	if n > 99 {
		return "", errors.New("invalind verse number")
	}
	sb := strings.Builder{}
	switch {
	case n > 1:
		fmt.Fprintf(&sb, "%d bottles of beer on the wall, %d bottles of beer.\n", n, n)
	case n == 1:
		sb.WriteString("1 bottle of beer on the wall, 1 bottle of beer.\n")
	case n == 0:
		sb.WriteString("No more bottles of beer on the wall, no more bottles of beer.\n")
	}

	switch {
	case n > 2:
		fmt.Fprintf(&sb, "Take one down and pass it around, %d bottles of beer on the wall.\n", n-1)
	case n == 2:
		sb.WriteString("Take one down and pass it around, 1 bottle of beer on the wall.\n")
	case n == 1:
		sb.WriteString("Take it down and pass it around, no more bottles of beer on the wall.\n")
	case n == 0:
		sb.WriteString("Go to the store and buy some more, 99 bottles of beer on the wall.\n")
	}
	return sb.String(), nil
}

// Verses returns verses of a given range.
func Verses(upperBound, lowerBound int) (string, error) {
	if lowerBound < 0 || upperBound > 99 || lowerBound > upperBound {
		return "", errors.New("invalid range")
	}
	sb := strings.Builder{}
	for i := upperBound; i >= lowerBound; i-- {
		v, err := Verse(i)
		if err != nil {
			return "", errors.New("erroneous verse")
		}
		fmt.Fprintf(&sb, "%s\n", v)
	}
	return sb.String(), nil
}

// Song recites the full beer song.
func Song() string {
	song, _ := Verses(99, 0)
	return song
}
