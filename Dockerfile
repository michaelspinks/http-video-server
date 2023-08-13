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

# Optional:
# To bind to a TCP port, runtime parameters must be supplied to the docker command.
# But we can document in the Dockerfile what ports
# the application is going to listen on by default.
# https://docs.docker.com/engine/reference/builder/#expose
EXPOSE 3000

# Set the PORT environment variable
ENV PORT=3000

WORKDIR /
COPY --from=build-env /app/http-video-server /http-video-server

# Run
CMD ["/http-video-server"]