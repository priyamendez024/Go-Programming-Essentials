# Chapter 21: Dockerfile
FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o app .

FROM alpine:3.17
WORKDIR /app
COPY --from=builder /app/app .
CMD ["./app"]
