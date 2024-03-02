FROM golang:1.17-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o csv-processor-api ./cmd/main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/csv-processor-api .

EXPOSE 8080

CMD ["./csv-processor-api"]
