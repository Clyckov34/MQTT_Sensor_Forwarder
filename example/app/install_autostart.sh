#!/bin/bash

set -e

SCRIPT_NAME="app"
CURRENT_DIR="$(cd "$(dirname "$0")" && pwd)"
SCRIPT_PATH="$CURRENT_DIR/$SCRIPT_NAME"

echo "========================================="
echo " Установка автозапуска app"
echo "========================================="

# ===== Проверка файла =====
if [ ! -f "$SCRIPT_PATH" ]; then
    echo "❌ Ошибка: файл $SCRIPT_NAME не найден!"
    echo "Ожидаемый путь: $SCRIPT_PATH"
    exit 1
fi

# ===== Делаем исполняемым =====
chmod +x "$SCRIPT_PATH"
echo "✅ Файл сделан исполняемым"

# ===== Получаем текущий crontab =====
CRON_TEMP=$(mktemp)

# Если cron пустой — не падаем
crontab -l 2>/dev/null > "$CRON_TEMP" || true

# ===== Удаляем старые записи wb8 =====
grep -v "$SCRIPT_PATH" "$CRON_TEMP" > "${CRON_TEMP}.clean" || true
mv "${CRON_TEMP}.clean" "$CRON_TEMP"

# ===== Добавляем новые =====
echo "50 23 * * * $SCRIPT_PATH" >> "$CRON_TEMP"

# ===== Применяем =====
crontab "$CRON_TEMP"
rm -f "$CRON_TEMP"

echo "✅ Автозапуск добавлен:"
echo "   - каждый день в 23:50"

# ===== Показ результата =====
echo ""
echo "📋 Текущий crontab:"
crontab -l

echo ""
echo "🎉 Установка завершена!"