package controller

import (
	"fmt"
	"net/http"
)

func LatestTemperature(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Latest temperature endpoint")
}

func MinTemperature(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Minimal temperature endpoint")
}

func MaxTemperature(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Maximal temperature endpoint")
}
