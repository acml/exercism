package scale

import "strings"

var chromaticScaleWithSharps = []string{"A", "A#", "B", "C", "C#", "D", "D#", "E", "F", "F#", "G", "G#"}
var chromaticScaleWithFlats = []string{"A", "Bb", "B", "C", "Db", "D", "Eb", "E", "F", "Gb", "G", "Ab"}

var flats = []string{"F", "Bb", "Eb", "Ab", "Db", "Gb", "d", "g", "c", "f", "bb", "eb"}

func useFlats(tonic string) bool {
	for _, v := range flats {
		if tonic == v {
			return true
		}
	}
	return false
}

// Scale given a tonic, or starting note, and a set of intervals, generates the
// musical scale starting with the tonic and following the specified interval
// pattern.
func Scale(tonic, interval string) []string {

	var scale []string
	if useFlats(tonic) {
		scale = chromaticScaleWithFlats
	} else {
		scale = chromaticScaleWithSharps
	}
	var i int
	var v string
	for i, v = range scale {
		if strings.EqualFold(v, tonic) {
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
			if !strings.EqualFold(tonic, scale[i]) {
				r[j] = scale[i]
				j++
			}
		}
		r = r[:j]
	}
	return r
}
