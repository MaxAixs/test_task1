FROM golang:1.22-alpine AS builder

# Устанавливаем рабочую директорию на cmd, где находится main.go
WORKDIR /app

# Копируем все файлы из проекта в контейнер, включая конфиг
COPY . .

# Запускаем go mod tidy и build
RUN go mod tidy
RUN go build -o /app/app cmd/main.go

# Этап с производственным образом
FROM alpine:latest
WORKDIR /root/

# Копируем скомпилированное приложение из builder
COPY --from=builder /app/app .

# Добавляем конфигурационный файл в нужную папку
COPY internal/config/config.yml /root/internal/config/config.yml

RUN chmod +x ./app
EXPOSE 8080
CMD ["./app"]
