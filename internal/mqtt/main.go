package mqtt

import (
	"MQTT/internal/config"
	"log"
	"strconv"
	"sync"
	"time"

	mt "github.com/eclipse/paho.mqtt.golang"
)

type indication map[string]float64

var (
	topics = make(indication)
	resMu  sync.RWMutex
)

func RunApp(s *config.Params) (indication, error) {
	clientOpt, err := newClient(s)
	if err != nil {
		return nil, err
	}

	client := mt.NewClient(clientOpt)
	if token := client.Connect(); token.WaitTimeout(10*time.Second) && token.Error() != nil {
		return nil, token.Error()
	}
	defer client.Disconnect(250)

	filters, err := getTopik(s.MqttTopicFile)
	if err != nil {
		return nil, err
	}

	expectedTopics := len(filters)
	received := 0

	// Канал для сигналов завершения (буфер 2: на таймаут и на все сообщения)
	done := make(chan bool, 2)

	// Колбэк для обработки сообщений
	messageHandler := func(client mt.Client, msg mt.Message) {
		defer func() {
			if r := recover(); r != nil {
				log.Printf("Recovered in message: %v\n", r)
			}
		}()

		resMu.Lock()
		defer resMu.Unlock()

		// Считаем только первое сообщение по топику
		if _, exists := topics[msg.Topic()]; !exists {
			received++
		}

		payLoadFloat, err := strconv.ParseFloat(string(msg.Payload()), 64)
		if err != nil {
			log.Println(err)
		}

		topics[msg.Topic()] = float64(payLoadFloat)

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

	return getIndication(), nil
}

// getIndication получаем готовые топики с данными
func getIndication() indication {
	resMu.RLock()
	defer resMu.RUnlock()

	result := make(indication, len(topics))
	for key, value := range topics {
		result[key] = value
	}
	return result
}
