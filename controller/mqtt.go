package controller

import (
	"fmt"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func ConsumeMQTT(client mqtt.Client, message mqtt.Message) {
	payload := message.Payload()
	fmt.Println(string(payload))
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
