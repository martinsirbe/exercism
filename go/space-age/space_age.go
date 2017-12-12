/*
Package space implements planet's age calculation.
*/
package space

const (
	EarthYearInSeconds = 31557600
)

// Age calculates given seconds as planet's age
func Age(seconds float64, planet string) float64 {
	planets := map[string]float64{
		"Earth":   EarthYearInSeconds,
		"Mercury": 0.2408467 * EarthYearInSeconds,
		"Venus":   0.61519726 * EarthYearInSeconds,
		"Mars":    1.8808158 * EarthYearInSeconds,
		"Jupiter": 11.862615 * EarthYearInSeconds,
		"Saturn":  29.447498 * EarthYearInSeconds,
		"Uranus":  84.016846 * EarthYearInSeconds,
		"Neptune": 164.79132 * EarthYearInSeconds,
	}

	return seconds / planets[planet]
}
