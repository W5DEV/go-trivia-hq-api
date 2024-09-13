# go-hp-trivia-api

The go-hp-trivia-api is the backend which hosts the questions and answers for my Harry Potter Trivia app, built to be hosted on local hardware utilizing Docker. It is built using Go, Gin and Gorm.

## To-Do

## Getting Started

Setup a PostgreSQL database and add the appropriate environment variables to the `app.env` file. Follow the `app.env.example` file for guidance.

Next, run the following commands to start the server:

```bash
go mod init github.com/W5DEV/go-hp-trivia-api

go run migrate/migrate.go

air
```

## Build Docker Image

To build the Docker image, run the following commands:

```bash
docker build --tag go-hp-trivia-api .

docker tag go-hp-trivia-api johnmwi/go-hp-trivia-api:latest

docker push johnmwi/go-hp-trivia-api:latest 
```
