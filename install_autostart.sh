#!/bin/bash
set -euo pipefail

SCRIPT_NAME="app"
CURRENT_DIR="$(cd "$(dirname "$0")" && pwd)"
SCRIPT_PATH="$CURRENT_DIR/$SCRIPT_NAME"

# Cron-строка: каждый час в 0 минут, с логированием
CRON_LINE="0 * * * * cd $CURRENT_DIR && ./$SCRIPT_NAME >> $CURRENT_DIR/cron.log 2>&1"
CRON_TEMP="$(mktemp)"

# Функция для очистки временных файлов при выходе
cleanup() {
    rm -f "$CRON_TEMP"
}
trap cleanup EXIT

echo "🔎 Диагностика перед установкой:"
echo "  Скрипт: $SCRIPT_PATH"
echo "   Текущая директория: $CURRENT_DIR"
echo "   Cron-строка: $CRON_LINE"

# Проверка файла
if [ ! -f "$SCRIPT_PATH" ]; then
    echo "❌ Файл не найден: $SCRIPT_PATH" >&2
    exit 1
fi

# Проверка прав на выполнение
if [ ! -x "$SCRIPT_PATH" ]; then
    echo "🔧 Добавляем права на выполнение: $SCRIPT_PATH"
    chmod +x "$SCRIPT_PATH"
fi

# Получаем текущий crontab (если есть)
if ! crontab -l 2>/dev/null > "$CRON_TEMP"; then
    echo "⚠️  Crontab пуст или недоступен, создаём новый" >&2
    > "$CRON_TEMP"  # Создаём пустой файл
fi

echo "📝 Текущий crontab:"
cat "$CRON_TEMP" || echo "(пусто)"

# Удаляем старую запись (строго как строку, без regex)
grep -vF "$SCRIPT_PATH" "$CRON_TEMP" > "${CRON_TEMP}.clean"
mv "${CRON_TEMP}.clean" "$CRON_TEMP"

# Добавляем новую строку
echo "$CRON_LINE" >> "$CRON_TEMP"

echo "📝 Итоговый crontab перед установкой:"
cat "$CRON_TEMP"

# Применяем новый crontab
if ! crontab "$CRON_TEMP"; then
    echo "❌ Не удалось установить crontab" >&2
    echo "Проверьте содержимое временного файла: $CRON_TEMP" >&2
    cat "$CRON_TEMP" >&2
    exit 1
fi

# Проверяем, что задание добавилось — ищем часть строки, которая точно есть в cron
if crontab -l | grep -qF "./$SCRIPT_NAME"; then
    echo "✅ Готово: запуск каждый час (в 0 минут)"
    echo "   Расписание: $CRON_LINE"
    echo "   Лог: $CURRENT_DIR/cron.log"
else
    echo "⚠️  Задание не было добавлено в crontab" >&2
    echo "📝 Содержимое crontab после установки:" >&2
    crontab -l >&2
    exit 1
fi
