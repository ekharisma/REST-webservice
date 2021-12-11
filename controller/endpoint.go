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
	timestamp   string
}

var id int = 0

func buildPayload(temperature int) string {
	payload := Payload{
		Id:          id,
		Temperature: temperature,
		timestamp:   time.Now().String(),
	}
	respond, err := json.MarshalIndent(payload, "", " ")
	if err != nil {
		panic(err.Error())
	}
	id += 1
	return string(respond)
}

func LatestTemperature(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s", buildPayload(latestTemperature))
}

func MinTemperature(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s", buildPayload(minTemperature))
}

func MaxTemperature(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s", buildPayload(maxTemperature))
}
