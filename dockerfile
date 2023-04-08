FROM alpine:latest
WORKDIR /app

COPY go.* ./

RUN go mod download

COPY *.go ./

EXPOSE 8080