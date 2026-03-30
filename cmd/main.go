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

	infoData(clientSensor)

	// Отправляем данные на сервер
	status, err := mqtt.SendJson(clientSensor)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("Status: " + status)
}

// infoData Выводит дынные в терминал чтобы посмотреть что пришло с датчиков
func infoData(c mqtt.Client) {
	for k, v := range c.SensorReadings {
		fmt.Printf("%v - %v\n", k, v)
	}

	fmt.Println("Констроллер ID:", c.ControllerID)
	fmt.Println("Email:", c.Email)
	fmt.Println("Токен:", c.Token)
}
