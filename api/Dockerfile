FROM golang:1.15-alpine as builder
WORKDIR /app
COPY . .
RUN go build -o api cmd/main/main.go

FROM alpine
WORKDIR /app
COPY --from=builder /app/api .
COPY config/local/server.yml config/docker/server.yml
ENV APP_ENV=docker

CMD ["./api", "--configDir", "config"]
