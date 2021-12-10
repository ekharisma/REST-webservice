package main

import (
	"ekharisma/challenge/controller"
	"log"
	"net/http"
)

func handleEndpoint() {
	http.HandleFunc("/sensor/latest", controller.LatestTemperature)
	http.HandleFunc("/sensor/min", controller.MinTemperature)
	http.HandleFunc("/sensor/max", controller.MaxTemperature)
}

func main() {
	handleEndpoint()
	log.Fatal(http.ListenAndServe(":8080", nil))
}
