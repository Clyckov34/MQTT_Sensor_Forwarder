package config

import (
	"errors"
	"os"

	"github.com/go-playground/validator/v10"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Server ServerConfig `yaml:"server"`
	Mqtt   MqttConfig   `yaml:"mqtt"`
}

type ServerConfig struct {
	Ip           string `yaml:"ip" validate:"required"`
	Token        string `yaml:"token" validate:"required"`
	ControllerId int    `yaml:"controller_id" validate:"required"`
	Login        string `yaml:"login" validate:"required"`
}

type MqttConfig struct {
	Ip        string `yaml:"ip" validate:"required"`
	Port      int    `yaml:"port" validate:"required"`
	Login     string `yaml:"login"`
	Password  string `yaml:"password"`
	FileTopic string `yaml:"file_topic" validate:"required"`
}

// LoadFile Загружаем файл c конфигурацией
func LoadYamlFile(fileName string) (*Config, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return nil, errors.New("Не удалось загрузить файл " + fileName + "Error:" + err.Error())
	}

	var config Config
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, errors.New("Не удалось распарсить файл " + fileName + "Error:" + err.Error())
	}

	return &config, nil
}

// ValidateConfig проверка данных в файле
func (c *Config) ValidateConfig() error {
	validate := validator.New()

	if err := validate.Struct(c); err != nil {
		return errors.New("Ошибка валидации конфигурации: " + err.Error())
	}

	return nil
}
