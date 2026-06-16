FROM golang:1.26-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o bin/server ./cmd/server

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/bin/server .

EXPOSE 3000

CMD ["./server"]