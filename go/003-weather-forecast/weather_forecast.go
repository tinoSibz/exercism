// Package weather provides weather forecast for a given city and condition.
package weather

// CurrentCondition holds the current weather condition.
var CurrentCondition string

// CurrentLocation holds the current location.
var CurrentLocation string

// Forecast returns the weather forecast for a given city and condition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
