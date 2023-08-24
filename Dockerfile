# Этап, на котором выполняется сборка приложения
FROM golang:1.16-alpine as builder
WORKDIR /build
COPY go.mod .
RUN go mod download
COPY . .
RUN go build -o /main cmd/app/main.go
# Финальный этап, копируем собранное приложение
FROM alpine:3
COPY --from=builder main /bin/main
EXPOSE 8050
ENTRYPOINT ["/bin/main"]