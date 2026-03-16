package clientMQTT

import (
	"MQTT/pkg/env"
	"fmt"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

// New инциализация приложение
func New(s *env.Server) (*mqtt.ClientOptions, error) {
	opt := mqtt.NewClientOptions()
	opt.AddBroker(fmt.Sprintf("tcp://%v:%v", s.MqttURL, s.MqttPort))

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
