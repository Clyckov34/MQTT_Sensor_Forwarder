package app

import (
	"MQTT/internal/clientMQTT"
	"MQTT/pkg/env"
	"fmt"
	"sync"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type topic map[string]byte

var (
	Res   = make(map[string]string)
	resMu sync.RWMutex
)

func Run(s *env.Server) (map[string]string, error) {
	clientOpt, err := clientMQTT.New(s)
	if err != nil {
		return nil, err
	}

	client := mqtt.NewClient(clientOpt)
	if token := client.Connect(); token.WaitTimeout(10*time.Second) && token.Error() != nil {
		return nil, token.Error()
	}
	defer client.Disconnect(250)

	// Создаём map[string]byte для SubscribeMultiple
	filters := filter()

	token := client.SubscribeMultiple(filters, message)
	if token.WaitTimeout(10*time.Second) && token.Error() != nil {
		return nil, token.Error()
	}

	// Ждём 60 секунд или прерываемся по сигналу
	done := make(chan bool, 1)
	time.AfterFunc(5*time.Second, func() {
		done <- true
	})

	<-done

	// Отписываемся от топиков — передаём именно топики (ключи из filters)
	topics := make([]string, 0, len(filters))
	for t := range filters {
		topics = append(topics, t)
	}
	unsubToken := client.Unsubscribe(topics...)
	if unsubToken.WaitTimeout(5*time.Second) && unsubToken.Error() != nil {
		return nil, unsubToken.Error()
	}

	return Res, nil
}

func filter() topic {
	return topic{
		"/devices/sauna_heater_ssr/controls/tempSetpoint_ssr": 2,
		"/devices/wb-adc/controls/Vin":                        2,
	}
}

func message(client mqtt.Client, msg mqtt.Message) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered in message: %v\n", r)
		}
	}()

	resMu.Lock()
	defer resMu.Unlock()
	Res[msg.Topic()] = string(msg.Payload())
}

func getResults() map[string]string {
	resMu.RLock()
	defer resMu.RUnlock()

	result := make(map[string]string, len(Res))
	for k, v := range Res {
		result[k] = v
	}
	return result
}
