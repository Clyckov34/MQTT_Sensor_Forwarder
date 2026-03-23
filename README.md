<div>
    <center><h1>Документация</h1></center>
    <p>Скрипт подключается к серверу MQTT запрашивает данные датчиков которые указаны в файле <b>topic.json</b> и отправляет на указанный сервер</p>
</div>    
<div>
    <h2>Настройка скрипта ENV</h2>
    <h3>Настройка окружение:</h3>
    <ol>
        <li>Откройте файл <b>app.env</b></li>
        <li>Укажите требуемые параметры</li>
        <ul>
            <li><code>SERVER_URL</code> - Адрес сервера куда будут отправляться показание датчиков</li>
            <li><code>CONTROLLER_ID</code> - Индификатор констролера</li>
            <li><code>CLIENT_ID</code> - Индификатор клиента</li>
            <li><code>CLIENT_TOKEN</code> - Токен клиента</li>
            <li><code>MQTT_URL</code> - URL (IP) адрес MQTT cервера</li>
            <li><code>MQTT_PORT</code> - Порт MQTT сервера</li>
            <li><code>MQTT_TOPIC_FILE</code> - Путь к файлу JSON с topic</li>
            <li><code>MQTT_USERNAME</code> - Логин MQTT сервера <code>Дополнительное поле</code></li>
            <li><code>MQTT_PASSWORD</code> - Пароль MQTT сервера <code>Дополнительное поле</code></li>
            
        </ul>
    </ol>
</div>