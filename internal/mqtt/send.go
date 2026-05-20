package mqtt

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"time"
)

type Client struct {
	Server         string             `json:"server"`
	ClientID       int                `json:"client_id"`
	Token          string             `json:"token"`
	ControllerID   int                `json:"controller_id"`
	SensorReadings map[string]float64 `json:"sensor_readings"`
}

// SendJsonPOST оптравляет данные на сервер методом POST
func SendJsonPOST(c Client) (status string, err error) {
	reqBody, err := json.Marshal(c)
	if err != nil {
		return "", err
	}

	ctx, _ := context.WithTimeout(context.Background(), time.Duration(10*time.Second))
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.Server, bytes.NewBuffer(reqBody))
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	return resp.Status, nil
}
