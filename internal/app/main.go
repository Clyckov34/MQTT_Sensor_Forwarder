package app

import (
	"MQTT/internal/clientMQTT"
	"MQTT/pkg/env"
	"fmt"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type topic map[string]byte

var driver = "/devices/energy_constant/controls/value"

// Run запуск приложение
func Run(server *env.Server) error {
	clientOpt, err := clientMQTT.New(server.IP, server.Port)
	if err != nil {
		return err
	}

	client := mqtt.NewClient(clientOpt)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		return token.Error()
	}

	token := client.SubscribeMultiple(filter(), message)
	if token.Wait() && token.Error() != nil {
		return token.Error()
	}

	// for i := 0; i < 5; i++ {
	// 	text := fmt.Sprintf("this is msg #%d!", i)
	// 	token := client.Publish(driver, 0, false, text)
	// 	token.Wait()
	// }

	token = client.Unsubscribe(driver)
	if token.Wait() && token.Error() != nil {
		return err
	}

	time.Sleep(10 * time.Second)

	client.Disconnect(250)
	time.Sleep(1 * time.Second)
	return nil
}

// filter фильтр топиков
func filter() topic {
	return topic{
		driver: 2,
	}
}

// message сообщение
func message(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("TOPIC: %s\n", msg.Topic())
	fmt.Printf("MSG: %s\n", msg.Payload())
}
