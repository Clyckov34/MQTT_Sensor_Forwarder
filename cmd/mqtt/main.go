package main

import (
	"MQTT/internal/clientMQTT"
	"MQTT/internal/config"
	"fmt"

	"log"
	"os"
)

var params *config.Server

func init() {
	if err := config.LoadFile("./app.env"); err != nil {
		log.Fatalln(err)
	}

	params = &config.Server{
		MqttURL:      os.Getenv("MQTT_URL"),
		MqttPort:     os.Getenv("MQTT_PORT"),
		MqttUserName: os.Getenv("MQTT_USERNAME"),
		MqttPassword: os.Getenv("MQTT_PASSWORD"),
		ClientID:     os.Getenv("CLIENT_ID"),
		ClientToken:  os.Getenv("CLIENT_TOKEN"),
	}

	err := config.CheckFile(params)
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
	topiks, err := clientMQTT.RunApp(params)
	if err != nil {
		log.Fatalln(err)
	}

	for k, v := range topiks {
		fmt.Println(k + " - " + v)
	}
}
