package main

import (
	"MQTT/internal/config"
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
		log.Fatalln(err)
	}

	//fmt.Println(clientSensor)

	// Отправляем данные на сервер
	statusCode, err := mqtt.SendJsonPOST(clientSensor)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("HTTP Status: ", statusCode)
}
