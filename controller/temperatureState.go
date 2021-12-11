package controller

var minTemperature int = 25
var maxTemperature int = 25
var latestTemperature int

func SetLatestTemperature(temperature int) {
	latestTemperature = temperature
	if minTemperature > temperature {
		minTemperature = temperature
	}
	if maxTemperature < temperature {
		maxTemperature = temperature
	}
}
