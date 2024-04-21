# syntax=docker/dockerfile:1

ARG GO_VERSION=1.22

FROM golang:${GO_VERSION}-alpine

ENV GO_ENV production

WORKDIR /usr/src/backend

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd

EXPOSE 9090

CMD ["./main"]