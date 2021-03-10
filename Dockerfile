FROM golang:1.14.6 AS build-env
ENV GO111MODULE=on

WORKDIR /app/notifyer
COPY go.mod .
COPY go.sum .

RUN go mod download
COPY . .

RUN go build -o notifyer
RUN chmod +x notifyer

CMD ["notifyer", "--version"]
