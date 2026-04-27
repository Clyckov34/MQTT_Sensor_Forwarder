package logging

import (
	"log"
	"os"
)

// LogToFile записать логов в файл
func LogToFile(data any, prefix string) error {
	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	l := log.New(file, prefix+": ", log.Ldate|log.Ltime)
	l.Println(data)

	return nil
}
