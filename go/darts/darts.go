package darts

import "math"

// Score returns the earned points in a single toss of a Darts game.
func Score(x, y float64) int {
	if distance := math.Sqrt(x*x + y*y); distance > 10 {
		return 0
	} else if distance <= 10 && distance > 5 {
		return 1
	} else if distance <= 5 && distance > 1 {
		return 5
	}
	return 10
}
