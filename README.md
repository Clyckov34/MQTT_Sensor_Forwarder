 <div align="center">
  <h1>🚀 MQTT API Relay</h1>
  <p>Простая и эффективная утилита для передачи данных с MQTT-брокера на внешний HTTP API.</p>
</div>

![GitHub release](https://img.shields.io/github/v/release/Clyckov34/MQTT_Sensor_Forwarder)
![License](https://img.shields.io/badge/license-MIT-green)
![Go](https://img.shields.io/badge/Go-1.20+-00ADD8?logo=go)
![Status](https://img.shields.io/badge/status-active-success)
![Last Commit](https://img.shields.io/github/last-commit/Clyckov34/MQTT_Sensor_Forwarder)
![Stars](https://img.shields.io/github/stars/Clyckov34/MQTT_Sensor_Forwarder?style=social)

<div>
    <img src="example/image.png" width="100%">
</div>
<div>
    <h2>Особенности:</h2>
    <ul>
        <li>Написана на языке Go, легковесная и производительная</li>
        <li>Поддержка многопоточности для параллельной обработки сообщений</li>
        <li>Не требует статического IP-адреса</li>
        <li>Легко настраиваемый автозагрузчик через cron</li>
        <li>Оптимизация для работы на платформе WirenBoard 6+</li>
        <li>Открытый исходный код и свободная лицензия</li>
    </ul>
</div>
<div>
    <h2>Зачем использовать:</h2>
    <ul>
        <li>Сбор и передача показаний датчиков в облачную инфраструктуру для умного дома и промышленного IoT</li>
        <li>Простота установки и настройки</li>
    </ul>
</div>
<div>
    <h2>📂 Структура проекта</h2>

```
├── config.env                  # Переменные окружения
├── topic.json                  # Список MQTT-топиков
├── app.log                     # Запись логов 
├── LICENSE                     # Лицензия
├── app                         # Бинарный файл приложения
└── install_autostart.sh        # Устанавливает автозапуск Cron
```

</div>
<div>
    <h2>📥 Вариант 1: Быстрый старт</h2>

```bash
wget https://github.com/Clyckov34/MQTT-API-Relay/releases/download/app-2.2.0/app.zip
unzip app.zip
cd app

# настроить конфиг
nano config.env

# настроить топик
nano topic.json

# Вариант 1: Запуск вручную
./app

# Вариант 2: Автозапуск
sudo ./install_autostart.sh

```

</div>
<div>
    <h2>🛠️ Вариант 2: Сборка</h2>

| Платформа          | Команда |
|--------------------|--------|
| Wiren Board (ARMv7) | `GOOS=linux GOARCH=arm GOARM=7 go build -o app cmd/main.go` |
| Linux (x64)        | `go build -o app cmd/main.go` |

</div>
<div>
    <h2>🔧 Настройка</h2>
    <h3>1. Настройка окружения</h3>
    <p>Откройте файл config.env и укажите параметры:</p>
    <ul>
        <li><code>SERVER</code> - Адрес сервера куда будут отправляться показания датчиков <code>string</code></li>
        <li><code>CONTROLLER_ID</code> - Идентификатор контроллера <code>integer</code></li>
        <li><code>CLIENT_ID</code> - Идентификатор клиента <code>integer</code></li>
        <li><code>CLIENT_TOKEN</code> - Токен клиента <code>string</code></li>
        <li><code>MQTT_SERVER</code> - URL (IP) адрес MQTT cервера <code>string</code></li>
        <li><code>MQTT_PORT</code> - Порт MQTT-сервера <code>integer</code></li>
        <li><code>MQTT_TOPIC_FILE</code> - Путь к файлу topic.json <code>string</code></li>
        <li><code>MQTT_USERNAME</code> - Логин MQTT-сервера <code>string</code> <code>Дополнительное поле</code></li>
        <li><code>MQTT_PASSWORD</code> - Пароль MQTT-сервера <code>string</code> <code>Дополнительное поле</code></li> 
    </ul>
</div>
<div>
<p>Пример файла окружения config.env</p>

```code
SERVER = "https://my_server_cloud.ru/post"
CONTROLLER_ID = 121

CLIENT_ID = 3241234
CLIENT_TOKEN = "qWeRtYuIoPaDfGHhkfwelwfk"

# MQTT
MQTT_SERVER = "localhost"
MQTT_PORT = 1883
MQTT_USERNAME = ""
MQTT_PASSWORD = ""
MQTT_TOPIC_FILE = "./topic.json"
```

</div>
<div>
    <h3>2. Настройка топиков</h3>
    <p>Файл topic.json содержит список топиков по которым будет подписываться:</p>

```json
{
  "topics": [
    {
      "path": "/devices/hwmon/controls/Board Temperature",
      "level_qos": 2
    },
    {
      "path": "/devices/hwmon/controls/CPU Temperature",
      "level_qos": 2
    }
  ]
}
```

</div>
<div>
    <h2>📡 QoS уровни</h2>

| Уровень   | Описание                                      |
| --------- | --------------------------------------------- |
| **QoS 0** | Максимум один раз (без гарантии доставки)     |
| **QoS 1** | Минимум один раз (возможны дубликаты)         |
| **QoS 2** | Ровно один раз (самый надёжный, но медленный) |

</div>
<div>
    <h2>📤 Формат отправляемых данных</h2>

```json
{
  "server": "https://httpbin.org/post",
  "client_id": "244235",
  "token": "Wefefor34rmcfree22svFFE",
  "controller_id": "000001",
  "sensor_readings": {
    "/devices/hwmon/controls/Board Temperature": 39.25,
    "/devices/hwmon/controls/CPU Temperature": 66.835,
    "/devices/sauna_floor_thermostat/controls/temperature": 31.9,
    "/devices/sauna_heater/controls/tempCurrent": 90.375,
    "/devices/sauna_heater_ssr/controls/tempSetpoint_ssr": 95,
    "/devices/wb-adc/controls/Vin": 50.26,
    "/devices/wb-m1w2_34/controls/External_Sensor_1": 13.3125,
    "/devices/wb-m1w2_34/controls/External_Sensor_2": 90.375,
    "/devices/wb-mr6cu_85/controls/MCU Temperature": 42.8,
    "/devices/wb-w1/controls/28-0000102149e4": 31.75,
    "/devices/wb-w1/controls/28-00001021f4a9": 32.187
  }
}

```

<h3>📌 Описание полей</h3>
<ul>
    <li><code>server</code> — адрес API сервера</li>
    <li><code>client_id</code> — идентификатор клиента</li>
    <li><code>token</code> — токен авторизации</li>
    <li><code>controller_id</code> — идентификатор устройства</li>
    <li><code>sensor_readings</code> — объект с данными датчиков
         <ul>
            <li>ключ — MQTT-топик</li>
            <li>значение — текущее значение датчика</li>
        </ul>
    </li> 
</ul>

</div>
<div>
    <h2>📜 Лицензия</h2>
    <p>MIT License</p>
</div>
