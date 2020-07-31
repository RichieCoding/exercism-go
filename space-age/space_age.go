// Package space allows for you to convert an age in seconds and calculate how old someone would be on a specific planet
package space

// Planet is name of the Planets
type Planet string

// Age function converts an age in seconds to their age on a choosen planet
func Age(s float64, p Planet) float64 {
	orbitalPeriod := map[Planet]float64{
		"Mercury": 0.2408467,
		"Venus":  0.61519726,
		"Earth": 1.0,
		"Mars": 1.8808158,
		"Jupiter": 11.862615,
		"Saturn": 29.447498,
		"Uranus": 84.016846,
		"Neptune": 164.79132,
	}

	return (s / 31557600) / orbitalPeriod[p]
}