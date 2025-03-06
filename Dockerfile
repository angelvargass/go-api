FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o go-api .

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/go-api .

EXPOSE 8080

CMD ["./go-api"]