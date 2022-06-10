# syntax=docker/dockerfile:1

FROM golang:1.16-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
# COPY *.go ./
COPY . ./

RUN go mod download
RUN go get github.com/gin-gonic/gin/binding@v1.7.7

RUN go build -o /docker-gs-ping

EXPOSE 8080

CMD [ "/docker-gs-ping" ]