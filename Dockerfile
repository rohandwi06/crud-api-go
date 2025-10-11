# Stage 1 (Builder)
FROM golang:1.25.0 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /app/main .

# Stage 2 (Final Image)
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/main .

COPY .env .  

EXPOSE 8080

# Jalankan aplikasi
CMD ["./main"]