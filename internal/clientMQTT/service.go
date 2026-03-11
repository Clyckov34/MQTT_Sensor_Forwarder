package clientMQTT

import (
	"fmt"
	"log"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)



// New инциализация приложение
func New(ip string, port int) (*mqtt.ClientOptions, error) {
	opt := mqtt.NewClientOptions()
	opt.AddBroker(fmt.Sprintf("tcp://%v:%v", ip, port))

	opt.SetClientID("avtomatika_MQT")
	opt.SetKeepAlive(2 * time.Second)
	opt.SetPingTimeout(1 * time.Second)

	opt.SetConnectionNotificationHandler(func(c mqtt.Client, cn mqtt.ConnectionNotification) {
		switch n := cn.(type) {
		case mqtt.ConnectionNotificationConnected:
			log.Println("[УВЕДОМЛЕНИЕ] подключение установлено")
		case mqtt.ConnectionNotificationConnecting:
			log.Printf("[УВЕДОМЛЕНИЕ] выполняется подключение (повторное=%t) [%d]\n", n.IsReconnect, n.Attempt)
		case mqtt.ConnectionNotificationFailed:
			log.Printf("[УВЕДОМЛЕНИЕ] ошибка подключения: %v\n", n.Reason)
		case mqtt.ConnectionNotificationLost:
			log.Printf("[УВЕДОМЛЕНИЕ] соединение потеряно: %v\n", n.Reason)
		case mqtt.ConnectionNotificationBroker:
			log.Printf("[УВЕДОМЛЕНИЕ] подключение к брокеру: %s\n", n.Broker.String())
		case mqtt.ConnectionNotificationBrokerFailed:
			log.Printf("[УВЕДОМЛЕНИЕ] ошибка подключения к брокеру: %v [%s]\n", n.Reason, n.Broker.String())
		}
	})

	return opt, nil
}


