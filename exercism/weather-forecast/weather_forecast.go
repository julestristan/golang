// Package weather provides tools to know the weather condition for a location.
package weather

// CurrentCondition represents the weather condition.
var CurrentCondition string

// CurrentLocation represents the city name.
var CurrentLocation string

// Forecast returns the location + the current weather in a city.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
