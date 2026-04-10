package mqtt

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Client struct {
	ServerUrl      string
	ID             int
	Token          string
	ControllerID   int
	SensorReadings map[string]float64
}

// SendJsonPOST оптравляет данные на сервер методом POST
func SendJsonPOST(c Client) (statusCode int, err error) {
	reqBody, err := json.Marshal(c)
	if err != nil {
		return 0, err
	}

	ctx, _ := context.WithTimeout(context.Background(), time.Duration(10*time.Second))
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.ServerUrl, bytes.NewBuffer(reqBody))
	if err != nil {
		return 0, err
	}

	req.Header.Set("Content-Type", "application/json")
	client := http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return 0, err
	}

	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return resp.StatusCode, fmt.Errorf("Сервер вернул статус ошибки: %d %s", resp.StatusCode, resp.Status)
	}

	return resp.StatusCode, nil
}
