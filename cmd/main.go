package main

import (
	"MQTT/internal/config"
	"MQTT/internal/mqtt"
	"fmt"

	"log"
	"os"
)

var params *config.Params

func init() {
	if err := config.LoadFile("./app.env"); err != nil {
		log.Fatalln(err)
	}

	params = &config.Params{
		ServerURL:     os.Getenv("SERVER_URL"),
		ControllerID:  os.Getenv("CONTROLLER_ID"),
		MqttURL:       os.Getenv("MQTT_URL"),
		MqttPort:      os.Getenv("MQTT_PORT"),
		MqttUserName:  os.Getenv("MQTT_USERNAME"),
		MqttPassword:  os.Getenv("MQTT_PASSWORD"),
		MqttTopicFile: os.Getenv("MQTT_TOPIC_FILE"),
		ClientEmail:   os.Getenv("CLIENT_EMAIL"),
		ClientToken:   os.Getenv("CLIENT_TOKEN"),
	}

	err := config.CheckParams(params)
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
	// Запрашиваем готовые топики с покозаниями
	clientSensor, err := mqtt.RunApp(params)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(clientSensor)

	// Отправляем данные на сервер
	status, err := mqtt.SendJsonPOST(clientSensor)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("Status: " + status)
}