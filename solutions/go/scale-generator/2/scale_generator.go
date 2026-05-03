package scale

import "strings"

// No Sharps or Flats:
// C major
// a minor

// Use Sharps:
// G, D, A, E, B, F# major
// e, b, f#, c#, g#, d# minor

// Use Flats:
// F, Bb, Eb, Ab, Db, Gb major
// d, g, c, f, bb, eb minor

var chromatic_scale_with_sharps = []string{"A", "A#", "B", "C", "C#", "D", "D#", "E", "F", "F#", "G", "G#"}
var chromatic_scale_with_flats = []string{"A", "Bb", "B", "C", "Db", "D", "Eb", "E", "F", "Gb", "G", "Ab"}

var sharps = []string{"G", "D", "A", "E", "B", "F#", "e", "b", "f#", "c#", "g#", "d#"}
var flats = []string{"F", "Bb", "Eb", "Ab", "Db", "Gb", "d", "g", "c", "f", "bb", "eb"}

func use_flats(tonic string) bool {
	for _, v := range flats {
		if tonic == v {
			return true
		}
	}
	return false
}

func Scale(tonic, interval string) []string {
	var i int
	var v string
	var scale []string

	if use_flats(tonic) {
		scale = chromatic_scale_with_flats
	} else {
		scale = chromatic_scale_with_sharps
	}
	for i, v = range scale {
		if strings.ToUpper(v) == strings.ToUpper(tonic) {
			break
		}
	}

	r := make([]string, len(scale))
	if interval == "" {
		for j := 0; j < len(scale); j++ {
			r[j] = scale[(i+j)%len(scale)]
		}
	} else {
		j := 0
		r[j] = scale[i]
		j++
		for _, v := range interval {
			if v == 'A' {
				i = (i + 3) % len(scale)
			} else if v == 'M' {
				i = (i + 2) % len(scale)
			} else {
				i = (i + 1) % len(scale)
			}
			if strings.ToUpper(tonic) != strings.ToUpper(scale[i]) {
				r[j] = scale[i]
				j++
			}
		}
		r = r[:j]
	}
	return r
}
