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

func RunApp(s *config.Config) (Client, error) {
	clientOpt, err := newClient(s)
	if err != nil {
		return Client{}, err
	}

	client := mt.NewClient(clientOpt)
	if token := client.Connect(); token.WaitTimeout(10*time.Second) && token.Error() != nil {
		return Client{}, token.Error()
	}
	defer client.Disconnect(250)

	topicPathAndQoS, err := getTopikFile(s.MqttTopicFile)
	if err != nil {
		return Client{}, err
	}

	// ЛОКАЛЬНОЕ состояние
	topics := make(indication)
	var resMu sync.RWMutex

	expectedTopics := len(topicPathAndQoS)
	received := 0

	done := make(chan struct{}, 1)

	messageHandler := func(client mt.Client, msg mt.Message) {
		defer func() {
			if r := recover(); r != nil {
				log.Printf("Recovered in message: %v\n", r)
			}
		}()

		topic := msg.Topic()

		payload, err := strconv.ParseFloat(string(msg.Payload()), 64)
		if err != nil {
			log.Println(err)
			return
		}

		resMu.Lock()
		defer resMu.Unlock()

		if _, exists := topics[topic]; !exists {
			received++
		}

		topics[topic] = payload

		if received >= expectedTopics {
			select {
			case done <- struct{}{}:
			default:
			}
		}
	}

	token := client.SubscribeMultiple(topicPathAndQoS, messageHandler)
	if token.WaitTimeout(10*time.Second) && token.Error() != nil {
		return Client{}, token.Error()
	}

	time.AfterFunc(30*time.Second, func() {
		select {
		case done <- struct{}{}:
		default:
		}
	})

	<-done

	// отписка
	unsubTopics := make([]string, 0, len(topicPathAndQoS))
	for t := range topicPathAndQoS {
		unsubTopics = append(unsubTopics, t)
	}

	unsubToken := client.Unsubscribe(unsubTopics...)
	if unsubToken.WaitTimeout(5*time.Second) && unsubToken.Error() != nil {
		return Client{}, unsubToken.Error()
	}

	return buildClient(s, topics, &resMu), nil
}

func buildClient(s *config.Config, topics indication, mu *sync.RWMutex) Client {
	mu.RLock()
	defer mu.RUnlock()

	result := make(indication, len(topics))
	for k, v := range topics {
		result[k] = v
	}

	return Client{
		ServerUrl:      s.Server,
		ID:             s.ClientID,
		Token:          s.ClientToken,
		ControllerID:   s.ControllerID,
		SensorReadings: result,
	}
}
