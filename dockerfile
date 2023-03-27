FROM golang:1.20-alpine as example-rest-on-go

WORKDIR /app

COPY go.* ./

RUN go mod download

COPY *.go ./

EXPOSE 8080