package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Payload struct {
	Id          int
	Temperature int
	Timestamp   string
}

var id int = 0

func buildPayload(temperature int) string {
	payload := Payload{
		Id:          id,
		Temperature: temperature,
		Timestamp:   time.Now().String(),
	}
	respond, err := json.MarshalIndent(payload, "", " ")
	if err != nil {
		panic(err.Error())
	}
	id += 1
	return string(respond)
}

func enableCORS(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func LatestTemperature(w http.ResponseWriter, r *http.Request) {
	enableCORS(&w)
	fmt.Fprintf(w, "%s", buildPayload(latestTemperature))
}

func MinTemperature(w http.ResponseWriter, r *http.Request) {
	enableCORS(&w)
	fmt.Fprintf(w, "%s", buildPayload(minTemperature))
}

func MaxTemperature(w http.ResponseWriter, r *http.Request) {
	enableCORS(&w)
	fmt.Fprintf(w, "%s", buildPayload(maxTemperature))
}
