<div>
    <center><h1>🚀 MQTT Sensor Forwarder</h1></center>
    <p>Скрипт для получения данных с MQTT-брокера и отправки на внешний API.
Подписывается на указанные MQTT-топики, получает данные с датчиков и пересылает их на заданный сервер.</p>
</div>
<div>
    <h2>📌 Возможности</h2>
    <ul>
        <li>Подключение к MQTT-брокеру</li>
        <li>Подписка на список топиков из JSON-файла</li>
        <li>Поддержка QoS 0 / 1 / 2</li>
        <li>Отправка данных на HTTP-сервер</li>
        <li>Поддержка авторизации MQTT</li>
    </ul>
</div>
<div>
    <h2>📂 Структура проекта</h2>

```
├── app.env                     # Переменные окружения
├── topic.json                  # Список MQTT-топиков
├── wb8                         # Скрипт
└── install_autostart.sh        # Устанавливает автозапуск Cron
```

</div>
<div>
    <h2>📥 Скачать и подготовить проект</h2>

```bash
$ wget https://github.com/Clyckov34/MQTT_Sensor_Forwarder/releases/download/wb8/WB-8.zip
$ unzip WB-8.zip
$ cd WB-8
```

</div>
<div>
    <h2>🔧 Настройка</h2>
    <h3>1. Настройка окружения</h3>
    <p>Откройте файл app.env и укажите параметры:</p>
    <ul>
        <li><code>SERVER_URL</code> - Адрес сервера куда будут отправляться показание датчиков</li>
        <li><code>CONTROLLER_ID</code> - Индификатор констролера</li>
        <li><code>CLIENT_EMAIL</code> - Почта клиента</li>
        <li><code>CLIENT_TOKEN</code> - Токен клиента</li>
        <li><code>MQTT_URL</code> - URL (IP) адрес MQTT cервера</li>
        <li><code>MQTT_PORT</code> - Порт MQTT сервера</li>
        <li><code>MQTT_TOPIC_FILE</code> - Путь к файлу JSON с topic</li>
        <li><code>MQTT_USERNAME</code> - Логин MQTT сервера <code>Дополнительное поле</code></li>
        <li><code>MQTT_PASSWORD</code> - Пароль MQTT сервера <code>Дополнительное поле</code></li> 
    </ul>
</div>
<div>
    <h3>2. Настройка топиков</h3>
    <p>Файл topic.json содержит список топиков по которым будет подписываться:</p>

```json
{
  "topics": [
    {
      "path": "/devices/hwmon/controls/Board Temperature",
      "level_qos": 1
    },
    {
      "path": "/devices/hwmon/controls/CPU Temperature",
      "level_qos": 1
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
    <h2>📤 Пример вывода</h2>

```bash
/devices/wb-w1/controls/28-00001021f4a9 - 32.187
/devices/sauna_heater_ssr/controls/tempSetpoint_ssr - 95
/devices/wb-m1w2_34/controls/External Sensor_1 - 13.625
/devices/sauna_heater/controls/tempCurrent - 92.875
/devices/hwmon/controls/CPU Temperature - 66.218
Контроллер ID: 000001
Email: my_mail@yandex.ru
Токен: igiOA8Xuh0zEdOf5haeX7LOmp1R16AzB
2026/03/30 21:01:02 Status: 200 OK
```

</div>
<div>
    <h2>📜 Лицензия</h2>
    <p>MIT License — свободно используйте и модифицируйте.</p>
</div>