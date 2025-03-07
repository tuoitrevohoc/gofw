FROM golang:latest AS builder

WORKDIR /app

COPY . .

RUN go build -o main ./cmd/service/main.go

EXPOSE 8000

CMD ["./main"]

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/main /app/
COPY ./config /app/config

CMD ["/app/main"]