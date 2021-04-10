package beer

import (
	"errors"
	"fmt"
)

// Verse returns a specific verse of the beer song.
func Verse(n int) (string, error) {
	if n > 99 {
		return "", errors.New("invalind verse number")
	}
	switch {
	case n > 2:
		return fmt.Sprintf(`%d bottles of beer on the wall, %d bottles of beer.
Take one down and pass it around, %d bottles of beer on the wall.
`, n, n, n-1), nil
	case n == 2:
		return `2 bottles of beer on the wall, 2 bottles of beer.
Take one down and pass it around, 1 bottle of beer on the wall.
`, nil
	case n == 1:
		return `1 bottle of beer on the wall, 1 bottle of beer.
Take it down and pass it around, no more bottles of beer on the wall.
`, nil
	default:
		return `No more bottles of beer on the wall, no more bottles of beer.
Go to the store and buy some more, 99 bottles of beer on the wall.
`, nil
	}
}

// Verses returns verses of a given range.
func Verses(upperBound, lowerBound int) (string, error) {
	if lowerBound < 0 || upperBound > 99 || lowerBound > upperBound {
		return "", errors.New("invalid range")
	}
	verses := ""
	for i := upperBound; i >= lowerBound; i-- {
		v, err := Verse(i)
		if err != nil {
			return "", errors.New("erroneous verse")
		}
		verses += v + "\n"
	}
	return verses, nil
}

// Song recites the full beer song.
func Song() string {
	song, _ := Verses(99, 0)
	return song
}
