# syntax=docker/dockerfile:1

FROM golang:1.16-alpine

WORKDIR /app

ENV GIN_MODE=release

COPY go.mod ./
COPY go.sum ./
# COPY *.go ./
COPY . ./

RUN go mod download


RUN go build -o /docker-gs-ping

EXPOSE 8080

CMD [ "/docker-gs-ping" ]