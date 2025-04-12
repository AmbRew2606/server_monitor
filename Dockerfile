FROM golang:1.23.5-alpine

WORKDIR /app

# Копируем go.mod и go.sum
COPY go.mod go.sum ./
RUN go mod download

# Копируем все файлы проекта
COPY . ./

# Сборка бинарника для мониторинга
RUN go build -o monitor ./cmd/monitor

# Устанавливаем права на выполнение для исполнимого файла
RUN chmod +x /app/monitor

# Команда, которая будет запускаться при старте контейнера
CMD ["./monitor"]
