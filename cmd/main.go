package main

import (
	"MQTT/internal/config"
	"MQTT/internal/logging"
	"MQTT/internal/mqtt"

	"log"
)

var params *config.Config

func init() {
	pr, err := config.LoadEnvFile("./config.env")
	if err != nil {
		log.Fatalln(err)
	}

	if err := pr.ValidateConfig(); err != nil {
		log.Fatalln(err)
	}

	params = pr
}

func main() {
	// Запрашиваем готовые топики с покозаниями
	clientSensor, err := mqtt.RunApp(params)
	if err != nil {
		logging.LogToFile(err, "ERROR MQTT: ")
		log.Fatalln(err)
	}

	// Отправляем данные на сервер
	statusCode, err := mqtt.SendJsonPOST(clientSensor)
	if err != nil {
		logging.LogToFile(statusCode, "ERROR Server: ")
		log.Fatalln(err)
	}

	logging.LogToFile(statusCode, "OK Server: ")
}
