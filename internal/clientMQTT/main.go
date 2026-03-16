package clientMQTT

import (
	"MQTT/internal/clientMQTT/service"
	"MQTT/pkg/env"
	"fmt"
	"sync"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var (
	Res   = make(map[string]string)
	resMu sync.RWMutex
)

func RunApp(s *env.Server) (map[string]string, error) {
	clientOpt, err := service.NewClient(s)
	if err != nil {
		return nil, err
	}

	client := mqtt.NewClient(clientOpt)
	if token := client.Connect(); token.WaitTimeout(10*time.Second) && token.Error() != nil {
		return nil, token.Error()
	}
	defer client.Disconnect(250)

	filters := *service.GetTopik()
	expectedTopics := len(filters)
	received := 0

	// Канал для сигналов завершения (буфер 2: на таймаут и на все сообщения)
	done := make(chan bool, 2)

	// Колбэк для обработки сообщений
	messageHandler := func(client mqtt.Client, msg mqtt.Message) {
		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("Recovered in message: %v\n", r)
			}
		}()

		resMu.Lock()
		defer resMu.Unlock()

		// Считаем только первое сообщение по топику
		if _, exists := Res[msg.Topic()]; !exists {
			received++
		}
		Res[msg.Topic()] = string(msg.Payload())

		// Если получили все ожидаемые топики, сигнализируем о завершении
		if received >= expectedTopics {
			done <- true
		}
	}

	token := client.SubscribeMultiple(filters, messageHandler)
	if token.WaitTimeout(10*time.Second) && token.Error() != nil {
		return nil, token.Error()
	}

	// Таймаут: 30 секунд — если не все сообщения пришли, всё равно завершаем
	time.AfterFunc(30*time.Second, func() {
		done <- true
	})

	// Ждём любого сигнала: либо все сообщения получены, либо время вышло
	<-done

	// Отписываемся от топиков
	topics := make([]string, 0, len(filters))
	for t := range filters {
		topics = append(topics, t)
	}

	unsubToken := client.Unsubscribe(topics...)
	if unsubToken.WaitTimeout(5*time.Second) && unsubToken.Error() != nil {
		return nil, unsubToken.Error()
	}

	return getResults(), nil
}

// getResults получаем готовые топики с данными
func getResults() map[string]string {
	resMu.RLock()
	defer resMu.RUnlock()

	result := make(map[string]string, len(Res))
	for k, v := range Res {
		result[k] = v
	}
	return result
}
