package config

import (
	"errors"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Server        string
	ControllerID  int
	MqttServer    string
	MqttPort      int
	MqttUserName  string
	MqttPassword  string
	MqttTopicFile string
	ClientID      int
	ClientToken   string
}

// LoadFile Загружаем файл с окружением
func LoadEnvFile(fileName string) (*Config, error) {
	if err := godotenv.Load(fileName); err != nil {
		return nil, errors.New("Не удалось загрузить файл " + fileName + "Error:" + err.Error())
	}

	controllerID, err := strconv.ParseInt(os.Getenv("CONTROLLER_ID"), 0, 64)
	if err != nil {
		return nil, err
	}

	mqttPort, err := strconv.ParseInt(os.Getenv("MQTT_PORT"), 0, 64)
	if err != nil {
		return nil, err
	}

	clientID, err := strconv.ParseInt(os.Getenv("CLIENT_ID"), 0, 64)
	if err != nil {
		return nil, err
	}

	return &Config{
		Server:        os.Getenv("SERVER"),
		ControllerID:  int(controllerID),
		ClientID:      int(clientID),
		ClientToken:   os.Getenv("CLIENT_TOKEN"),
		MqttServer:    os.Getenv("MQTT_SERVER"),
		MqttPort:      int(mqttPort),
		MqttUserName:  os.Getenv("MQTT_USERNAME"),
		MqttPassword:  os.Getenv("MQTT_PASSWORD"),
		MqttTopicFile: os.Getenv("MQTT_TOPIC_FILE"),
	}, nil
}

// ValidateConfig проверка данных в файле
func (c *Config) ValidateConfig() error {
	if isEmpty(c.ClientID) {
		return errors.New("Не указан CLIENT_ID")
	} else if isEmpty(c.ClientToken) {
		return errors.New("Не указан CLIENT_TOKEN")
	} else if isEmpty(c.Server) {
		return errors.New("Не указан SERVER")
	} else if isEmpty(c.MqttServer) {
		return errors.New("Не указан MQTT_SERVER")
	} else if isEmpty(c.MqttPort) {
		return errors.New("Не указан MQTT_PORT")
	} else if isEmpty(c.ControllerID) {
		return errors.New("Не указан CONTROLLER_ID")
	} else if isEmpty(c.MqttTopicFile) {
		return errors.New("Не указан MQTT_TOPIC_FILE")
	} else {
		return nil
	}
}

// isEmpty проверка на пустоту
func isEmpty(data any) bool {
	switch v := data.(type) {
	case nil:
		return true
	case string:
		return v == ""
	case int, int8, int16, int32, int64:
		return v == 0
	case uint, uint8, uint16, uint32, uint64:
		return v == 0
	case float32, float64:
		return v == 0.0
	case bool:
		return !v
	default:
		return data == nil
	}
}
