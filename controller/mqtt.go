package controller

import (
	"fmt"
	"log"

	payload "ekharisma/challenge/protobuf"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/golang/protobuf/proto"
)

func parseMessage(message []byte) (int, int, string) {
	pl := &payload.Payload{}
	// var payload []byte
	if err := proto.Unmarshal(message, pl); err != nil {
		log.Fatal("Failed to parse message from MQTT: ", err)
	}
	return int(pl.Id), int(pl.Temperature), pl.Timestamp
}

func ConsumeMQTT(client mqtt.Client, message mqtt.Message) {
	id, temperature, timestamp := parseMessage(message.Payload())
	fmt.Println("id ", id)
	fmt.Println("temperature: ", temperature)
	fmt.Println("Received at: ", timestamp)
}

func MQTTInit(broker string) mqtt.Client {
	opts := mqtt.NewClientOptions()
	opts.AddBroker("tcp://" + broker)
	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	return client
}
