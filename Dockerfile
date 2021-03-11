# Build stage
FROM golang:1.14.6 AS build-env
ENV GO111MODULE=on
ENV CGO_ENABLED=0

WORKDIR /app/notifyer
COPY go.mod .
COPY go.sum .

RUN go mod download
COPY . .

RUN go build -o notifyer
RUN chmod +x notifyer

# Release stage
FROM alpine:latest AS release
WORKDIR /app/
COPY --from=build-env /app/notifyer .

CMD ["notifyer", "--version"]
