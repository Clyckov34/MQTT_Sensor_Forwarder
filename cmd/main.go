package main

import (
	"MQTT/internal/config"
	"MQTT/internal/mqtt"
	"fmt"

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

	// Отправляем данные на сервер
	statusCode, err := mqtt.SendJsonPOST(clientSensor)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("\nHTTP Status: %v\n", statusCode)
}
