# Используем официальный образ Go на базе Alpine как базовый
FROM golang:1.21-alpine AS builder

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем go.mod и go.sum файлы
COPY go.mod go.sum ./

# Загружаем зависимости
RUN go mod download

# Копируем весь исходный код
COPY . .

# Компилируем приложение
RUN go build -o main ./cmd/main.go

# Создаем финальный образ на базе Alpine
FROM alpine:latest

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем скомпилированное приложение из первого этапа
COPY --from=builder /app/main /app/main

# Копируем файл конфигурации, если он есть
COPY .env /app/.env

# Делаем скомпилированное приложение исполняемым
RUN chmod +x /app/main

# Открываем порт, на котором работает приложение
EXPOSE 8080

# Команда по умолчанию для запуска приложения
CMD ["/app/main"]
