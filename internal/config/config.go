package config

import (
	"errors"

	"github.com/joho/godotenv"
)

type Params struct {
	ServerURL     string
	ControllerID  string
	MqttURL       string
	MqttPort      string
	MqttUserName  string
	MqttPassword  string
	MqttTopicFile string
	ClientEmail   string
	ClientToken   string
}

// LoadFile Загружаем файл с окружением
func LoadFile(path string) error {
	if err := godotenv.Load(path); err != nil {
		return errors.New("Не удалось загрузить файл " + path + "Error:" + err.Error())
	}

	return nil
}

// CheckParams проверка данных в файле
func CheckParams(ser *Params) error {
	if !checkParam(ser.ClientEmail) {
		return errors.New("Не указан CLIENT_EMAIL")
	} else if !checkParam(ser.ClientToken) {
		return errors.New("Не указан CLIENT_TOKEN")
	} else if !checkParam(ser.ServerURL) {
		return errors.New("Не указан SERVER_URL")
	} else if !checkParam(ser.MqttURL) {
		return errors.New("Не указан MQTT_URL")
	} else if !checkParam(ser.MqttPort) {
		return errors.New("Не указан MQTT_PORT")
	} else if !checkParam(ser.ControllerID) {
		return errors.New("Не указан CONTROLLER_ID")
	} else if !checkParam(ser.MqttTopicFile) {
		return errors.New("Не указан MQTT_TOPIC_FILE")
	} else {
		return nil
	}
}

// checkParam проверка параметров
func checkParam(param string) bool {
	if len(param) == 0 {
		return false
	}

	return true
}
