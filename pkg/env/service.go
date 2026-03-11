package env

import (
	"errors"
	"strconv"

	"github.com/joho/godotenv"
)

type Server struct {
	IP   string
	Port int
}

// LoadFile Загружаем файл с окружением
func LoadFile(path string) error {
	if err := godotenv.Load(path); err != nil {
		return errors.New("Не удалось загрузить файл " + path + "Error:" + err.Error())
	}

	return nil
}

// CheckFile проверка данных в файле
func CheckFile(ip, port string) (*Server, error) {
	if len(ip) == 0 {
		return nil, errors.New("Неверный формат URL в файле .env")
	}

	portINT, err := strconv.Atoi(port)
	if len(port) == 0 || err != nil {
		return nil, errors.New("Неверный формат PORT в файле .env")
	}

	return &Server{
		IP:   ip,
		Port: portINT,
	}, nil
}
