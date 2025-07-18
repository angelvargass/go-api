FROM golang:1.24-alpine AS builder

WORKDIR /go/src/app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -o /go/bin/go-api ./cmd/main.go

FROM gcr.io/distroless/static-debian12

COPY --from=builder /go/bin/go-api /

EXPOSE 8080

CMD ["./go-api"]
