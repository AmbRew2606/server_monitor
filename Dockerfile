FROM golang:1.21-alpine

WORKDIR /app

# Копируем go.mod и go.sum, чтобы только при изменении зависимостей перезапускалась сборка
COPY go.mod go.sum ./
RUN go mod download

# Копируем все файлы проекта
COPY . .

# Сборка бинарника для мониторинга
RUN go build -o monitor ./cmd/monitor

# Сборка бинарника для Telegram бота
RUN go build -o telegrambot ./pkg/telegram

# Команда, которая будет запускаться при старте контейнера
CMD ["sh", "-c", "./monitor & ./telegrambot"]
