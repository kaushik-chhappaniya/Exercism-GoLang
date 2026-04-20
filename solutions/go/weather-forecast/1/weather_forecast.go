// Package weather provides tools forecasting the weather.
package weather

var (
	// CurrentCondition takes the current weather condition.
    CurrentCondition string 
	// CurrentLocation this shows the current location.
    CurrentLocation  string 
)

// Forecast function forecasts the weather based on city and condition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
