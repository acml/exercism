// Package pythagorean provides functions to find all Pythagorean triplets for
// which `a + b + c = N`.
package pythagorean

// Triplet is a set of three natural numbers, {a, b, c} which form a Pythagorean
// triplet
type Triplet [3]int

// Range returns a list of all Pythagorean triplets with sides in the range min
// to max inclusive.
func Range(min, max int) []Triplet {
	result := []Triplet{}

	if min > max || max-min < 2 {
		return result
	}

	for a := min; a < max-1; a++ {
		for b := a + 1; b < max; b++ {
			for c := b + 1; c <= max; c++ {
				if a*a+b*b == c*c {
					result = append(result, Triplet{a, b, c})
				}
			}
		}
	}
	return result
}

// Sum returns a list of all Pythagorean triplets where the sum a+b+c (the
// perimeter) is equal to p.
func Sum(p int) []Triplet {
	result := []Triplet{}
	triplets := Range(1, p/2)
	for _, triplet := range triplets {
		if triplet[0]+triplet[1]+triplet[2] == p {
			result = append(result, triplet)
		}
	}
	return result
}
