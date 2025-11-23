// Package weather document current weather
// based on condition and location.
package weather

var (
	// CurrentCondition represents weather condition.
	CurrentCondition string
	// CurrentLocation represents location.
	CurrentLocation string
)

// Forecast return string value that indicate
// the current weather condition with location.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
