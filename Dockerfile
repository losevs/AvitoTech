FROM golang:1.20.3

WORKDIR /usr/src/app

COPY . .
RUN go mod tidy