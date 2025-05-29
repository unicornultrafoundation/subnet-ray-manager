# syntax=docker/dockerfile:1

FROM golang:1.24 AS builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

RUN go build -o subnet-ray-manager ./cmd/subnet-ray-manager

FROM alpine:3.21

WORKDIR /app

COPY --from=builder /app/subnet-ray-manager ./

CMD ["./subnet-ray-manager"]
