package space

// Planet will be the name of the planet which we will be making calculations
// for.
type Planet string

var earthYear = map[Planet]float64{
	"Mercury": 0.2408467,
	"Venus":   0.61519726,
	"Earth":   1.0,
	"Mars":    1.8808158,
	"Jupiter": 11.862615,
	"Saturn":  29.447498,
	"Uranus":  84.016846,
	"Neptune": 164.79132,
}

// Age given an age in seconds, calculates how old someone would be on other
// planets.
func Age(seconds float64, planet Planet) float64 {

	return seconds / (31557600.0 * earthYear[planet])
}
