FROM golang:alpine AS builder

WORKDIR /go/src/app

COPY go.mod .
RUN go mod download

EXPOSE 8080

CMD ["go", "run", "server.go"]