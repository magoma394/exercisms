// Package weather provides functions to track weather conditions and locations.
package weather


var (
	// CurrentCondition represents the current weather condition.
	CurrentCondition string
	// CurrentLocation represents the current location.
	CurrentLocation  string
)

// Forecast returns a formatted weather forecast string for a given city and condition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
