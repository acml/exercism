package darts

import "math"

func Score(x, y float64) int {
	distance := math.Sqrt(x*x + y*y)
	if distance > 10 {
		return 0
	} else if distance <= 10 && distance > 5 {
		return 1
	} else if distance <= 5 && distance > 1 {
		return 5
	}
	return 10
}
