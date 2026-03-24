package mqtt

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"
)

// SendJson оптравляет данные на сервер
func SendJson(url string, data indication) (status string, err error) {
	reqBody, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	client := http.Client{
		Timeout: time.Duration(10 * time.Second),
	}

	resp, err := client.Post(url, "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	return resp.Status, nil
}
