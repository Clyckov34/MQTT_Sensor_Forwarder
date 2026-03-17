package config

import (
	"errors"

	"github.com/joho/godotenv"
)

type Server struct {
	MqttURL      string
	MqttPort     string
	MqttUserName string
	MqttPassword string
	ClientID     string
	ClientToken  string
}

// LoadFile Загружаем файл с окружением
func LoadFile(path string) error {
	if err := godotenv.Load(path); err != nil {
		return errors.New("Не удалось загрузить файл " + path + "Error:" + err.Error())
	}

	return nil
}

// CheckFile проверка данных в файле
func CheckFile(ser *Server) error {
	if len(ser.MqttPort) == 0 || len(ser.MqttURL) == 0 || len(ser.ClientID) == 0 || len(ser.ClientToken) == 0 {
		return errors.New("Не заполнены обязательные поля в app.env")
	}

	return nil
}
