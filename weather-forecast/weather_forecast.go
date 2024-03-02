
// Package weather forecasts the current weather condition of
// various cities in Goblinocus.
package weather

// CurrentCondition represents a string describing the
// current weather conditions.
var CurrentCondition string
// CurrentLocation represents a string describing the
// current location of the forecast.
var CurrentLocation string

// Forecast returns a string that that shows the current weather
// conditions at a particular lacation.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
