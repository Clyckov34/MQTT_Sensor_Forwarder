package mqtt

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"
)

type Client struct {
	ServerUrl      string
	Email          string
	Token          string
	ControllerID   string
	SensorReadings map[string]float64
}

// SendJsonPOST оптравляет данные на сервер методом POST
func SendJsonPOST(c Client) (status string, err error) {
	reqBody, err := json.Marshal(c)
	if err != nil {
		return "", err
	}

	client := http.Client{
		Timeout: time.Duration(10 * time.Second),
	}

	resp, err := client.Post(c.ServerUrl, "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	return resp.Status, nil
}
