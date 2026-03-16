package main

import (
	"MQTT/internal/app"
	"MQTT/pkg/env"
	"fmt"

	"log"
	"os"
)

var params *env.Server

func init() {
	if err := env.LoadFile("./app.env"); err != nil {
		log.Fatalln(err)
	}

	params = &env.Server{
		MqttURL:      os.Getenv("MQTT_URL"),
		MqttPort:     os.Getenv("MQTT_PORT"),
		MqttUserName: os.Getenv("MQTT_USERNAME"),
		MqttPassword: os.Getenv("MQTT_PASSWORD"),
		ClientID:     os.Getenv("CLIENT_ID"),
		ClientToken:  os.Getenv("CLIENT_TOKEN"),
	}

	err := env.CheckFile(params)
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
	res, err := app.Run(params)
	if err != nil {
		log.Fatalln(err)
	}

	for k, v := range res {
		fmt.Println(k + " - " + v)
	}
}
