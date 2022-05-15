# syntax=docker/dockerfile:1

FROM golang:1.16-alpine

WORKDIR /app

ENV GO_ENV=dev

COPY go.mod ./
COPY go.sum ./
# COPY *.go ./
COPY . ./

RUN go mod download


RUN go build -o /docker-gs-ping

EXPOSE 8080

CMD [ "/docker-gs-ping" ]