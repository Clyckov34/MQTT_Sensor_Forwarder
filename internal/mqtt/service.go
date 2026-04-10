package mqtt

import (
	"MQTT/internal/config"
	"encoding/json"
	"fmt"
	"os"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type topic map[string]byte

type TopicsData struct {
	TopicJSON2 []TopicJSON `json:"topics"`
}

type TopicJSON struct {
	Path     string `json:"path"`
	LevelQoS byte   `json:"level_qos"`
}

// NewClient инциализация приложение
func newClient(s *config.Config) (*mqtt.ClientOptions, error) {
	opt := mqtt.NewClientOptions()
	opt.AddBroker(fmt.Sprintf("tcp://%v:%v", s.MqttServer, s.MqttPort))

	opt.SetClientID("avtomatika_MQT")
	opt.SetKeepAlive(2 * time.Second)
	opt.SetPingTimeout(1 * time.Second)
	opt.SetCleanSession(true)

	if len(s.MqttUserName) > 0 && len(s.MqttPassword) > 0 {
		opt.SetUsername(s.MqttUserName)
		opt.SetPassword(s.MqttPassword)
	}

	return opt, nil
}

// getTopik получить топик
func getTopik(nameFile string) (topic, error) {
	byte, err := os.ReadFile(nameFile)
	if err != nil {
		return nil, err
	}

	var data TopicsData

	err = json.Unmarshal(byte, &data)
	if err != nil {
		return nil, err
	}

	topic := make(topic, 0)

	for _, v := range data.TopicJSON2 {
		topic[v.Path] = v.LevelQoS
	}

	return topic, nil
}
