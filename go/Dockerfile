FROM golang:1.21.5-alpine3.19

RUN apk add --no-cache git bash

RUN go install github.com/cosmtrek/air@latest

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

CMD [ "air" ]
