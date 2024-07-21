# syntax=docker/dockerfile:1

FROM golang:latest as builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY ./ ./
RUN go build -o go-hp-trivia-api .

EXPOSE 33500

CMD ["./go-hp-trivia-api"]