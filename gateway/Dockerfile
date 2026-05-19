FROM golang:1.25-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o gateway-binary ./gateway/cmd/api

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/gateway-binary .

EXPOSE 8080

CMD ["./gateway-binary"]