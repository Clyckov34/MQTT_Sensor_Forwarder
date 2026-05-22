#!/bin/bash
set -euo pipefail

SCRIPT_NAME="app"
CURRENT_DIR="$(cd "$(dirname "$0")" && pwd)"
SCRIPT_PATH="$CURRENT_DIR/$SCRIPT_NAME"

# Каждую минуту
CRON_LINE="* * * * * cd $CURRENT_DIR && ./$SCRIPT_NAME >> $CURRENT_DIR/cron.log 2>&1"

CRON_TEMP="$(mktemp)"
trap 'rm -f "$CRON_TEMP" "${CRON_TEMP}.clean"' EXIT

# Проверка файла
if [ ! -f "$SCRIPT_PATH" ]; then
    echo "❌ Файл не найден: $SCRIPT_PATH"
    exit 1
fi

# Права на выполнение
chmod +x "$SCRIPT_PATH"

# Получаем текущий crontab
crontab -l > "$CRON_TEMP" 2>/dev/null || true

echo "📋 Текущий crontab:"
cat "$CRON_TEMP" || true

# Удаляем старые записи app
grep -vF "./$SCRIPT_NAME" "$CRON_TEMP" > "${CRON_TEMP}.clean" || true
mv "${CRON_TEMP}.clean" "$CRON_TEMP"

# Добавляем новую запись
echo "$CRON_LINE" >> "$CRON_TEMP"

echo "📝 Новый crontab:"
cat "$CRON_TEMP"

# Устанавливаем
crontab "$CRON_TEMP"

echo "✅ Готово"
echo "📌 Проверка:"
crontab -l