#!/bin/bash
set -e

SCRIPT_NAME="app"
CURRENT_DIR="$(cd "$(dirname "$0")" && pwd)"
SCRIPT_PATH="$CURRENT_DIR/$SCRIPT_NAME"

CRON_LINE="0 * * * * $SCRIPT_PATH"
CRON_TEMP="$(mktemp)"

# Проверка файла
if [ ! -f "$SCRIPT_PATH" ]; then
    echo "❌ Файл не найден: $SCRIPT_PATH"
    exit 1
fi

# Делаем исполняемым
chmod +x "$SCRIPT_PATH"

# Получаем текущий crontab (если есть)
crontab -l 2>/dev/null > "$CRON_TEMP" || true

# Удаляем старую запись (строго как строку, без regex)
grep -vF "$SCRIPT_PATH" "$CRON_TEMP" > "${CRON_TEMP}.clean"
mv "${CRON_TEMP}.clean" "$CRON_TEMP"

# Добавляем новую строку
echo "$CRON_LINE" >> "$CRON_TEMP"

# Применяем
crontab "$CRON_TEMP"

# Очистка
rm -f "$CRON_TEMP"

echo "✅ Готово: запуск каждый час"