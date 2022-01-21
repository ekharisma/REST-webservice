package test

import (
	"ekharisma/challenge/controller"
	"reflect"
	"testing"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

const broker = "broker.hivemq.com:1883"

func IsInstanceOf(objectPtr, typePtr interface{}) bool {
	return reflect.TypeOf(objectPtr) == reflect.TypeOf(typePtr)
}

func TestMQTTClient(t *testing.T) {
	v := controller.MQTTInit(broker)
	opts := mqtt.NewClientOptions()
	opts.AddBroker("tcp://" + broker)
	if IsInstanceOf(v, mqtt.NewClient(opts)) {
		t.Error("Not a MQTT client")
	}
}
