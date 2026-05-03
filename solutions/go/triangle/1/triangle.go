// Package triangle should have a package comment that summarizes what it's about.
package triangle

import "math"

// Kind is what triangle type we have
type Kind int

// Possible triangle types
const (
	NaT Kind = iota // not a triangle
	Equ             // equilateral
	Iso             // isosceles
	Sca             // scalene
)

// isTriangle determines if any given three length matches a possible triangle
// or not.
func isTriangle(a, b, c float64) bool {
	if math.IsNaN(a) || math.IsNaN(b) || math.IsNaN(c) {
		return false
	}

	if math.IsInf(a, 0) || math.IsInf(b, 0) || math.IsInf(c, 0) {
		return false
	}

	if a <= 0 || b <= 0 || c <= 0 {
		return false
	}

	if a+b < c || b+c < a || c+a < b {
		return false
	}

	return true
}

// KindFromSides gives the triangle type.
func KindFromSides(a, b, c float64) Kind {
	var k Kind

	if !isTriangle(a, b, c) {
		k = NaT
	} else {
		if a == b && b == c {
			k = Equ
		} else if a == b || b == c || c == a {
			k = Iso
		} else {
			k = Sca
		}
	}

	return k
}
