package main

import (
	"ekharisma/challenge/controller"
	"log"
	"net/http"
)

const broker = "broker.hivemq.com:1883"

const topic = "iot/thermostat"

// var message string = ""

func handleFunc() {
	http.HandleFunc("/sensor/latest", controller.LatestTemperature)
	http.HandleFunc("/sensor/min", controller.MinTemperature)
	http.HandleFunc("/sensor/max", controller.MaxTemperature)
}

func main() {
	handleFunc()
	// go consumeMQTT()
	client := controller.MQTTInit(broker)
	if token := client.Subscribe(topic, 0, controller.ConsumeMQTT); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	log.Fatal(http.ListenAndServe(":8000", nil))
}
