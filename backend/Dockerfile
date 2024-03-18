# syntax=docker/dockerfile:1

ARG GO_VERSION=1.22
FROM golang:${GO_VERSION}-alpine
ENV GO_ENV production
WORKDIR /go/src/app
# Download Go modules
COPY go.mod go.sum ./
RUN go mod download
COPY . .
# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /image-converter-backend
# To bind to a TCP port, runtime parameters must be supplied to the docker command.
# But we can (optionally) document in the Dockerfile what ports
# the application is going to listen on by default.
# https://docs.docker.com/engine/reference/builder/#expose
EXPOSE 8080
# Run
CMD [ "/image-converter-backend" ]