# syntax=docker/dockerfile:1

FROM golang:1.20.4 AS build-env

# Set destination for COPY
WORKDIR /app

COPY go.mod ./
run go mod download

# Build
COPY . .
RUN go build -o http-video-server

FROM debian:bullseye-slim

WORKDIR /
COPY --from=build-env /app/http-video-server /http-video-server

# Run
CMD ["/http-video-server"]