package clientMQTT

import (
	"MQTT/internal/config"
	"fmt"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type topic map[string]byte

// NewClient инциализация приложение
func newClient(s *config.Server) (*mqtt.ClientOptions, error) {
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

// getTopik получить топик
func getTopik() topic {
	return topic{
		"/devices/sauna_heater_ssr/controls/tempSetpoint_ssr":  2,
		"/devices/wb-adc/controls/Vin":                         2,
		"/devices/hwmon/controls/Board Temperature":            2,
		"/devices/hwmon/controls/CPU Temperature":              2,
		"/devices/sauna_floor_thermostat/controls/temperature": 2,
		"/devices/sauna_heater/controls/tempCurrent":           2,
		"/devices/wb-m1w2_34/controls/External_Sensor_1":       2,
		"/devices/wb-m1w2_34/controls/External_Sensor_2":       2,
		"/devices/wb-mr6cu_85/controls/MCU Temperature":        2,
		"/devices/wb-w1/controls/28-0000102149e4":              2,
		"/devices/wb-w1/controls/28-00001021f4a9":              2,
	}
}
