# Используем официальный образ Go как базовый для сборки
FROM golang:1.23 AS builder

# Устанавливаем рабочую директорию внутри контейнера
WORKDIR /app

# Копируем go.mod и go.sum для загрузки зависимостей
COPY go.mod go.sum ./
RUN go mod download

# Копируем весь проект
COPY . .

# Собираем приложение
RUN CGO_ENABLED=0 GOOS=linux go build -o /blog-graphql cmd/blog/main.go

# Используем минимальный образ для выполнения
FROM alpine:latest

# Устанавливаем зависимости для работы приложения (если нужны)
RUN apk --no-cache add ca-certificates

# Устанавливаем рабочую директорию
WORKDIR /root/

# Копируем скомпилированный бинарник из builder
COPY --from=builder /my-blog-graphql .

# Копируем миграции (если они нужны внутри контейнера)
COPY migrations /root/migrations

# Указываем порт, который будет слушать приложение
EXPOSE 8080

# Команда для запуска приложения
CMD ["./my-blog-graphql"]