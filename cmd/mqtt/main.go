package main

import (
	"MQTT/internal/app"
	"MQTT/pkg/env"

	"log"
	"os"
)

var params *env.Server

func init() {
	if err := env.LoadFile("./app.env"); err != nil {
		log.Fatalln(err)
	}

	res, err := env.CheckFile(os.Getenv("MQTT_URL"), os.Getenv("MQTT_PORT"))
	if err != nil {
		log.Fatalln(err)
	}

	params = res
}

func main() {
	if err := app.Run(params); err != nil {
		log.Fatalln(err)
	}
}
