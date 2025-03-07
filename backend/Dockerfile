FROM golang:latest

WORKDIR /app

COPY . .

RUN go build -o main ./cmd/service/main.go

EXPOSE 8000

CMD ["./main"]