FROM golang:1.24.4-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o wordbook-generator .

FROM alpine:3.20

WORKDIR /app

COPY --from=builder /app/wordbook-generator .
COPY --from=builder /app/frontend /app/frontend
COPY --from=builder /app/resources /app/resources

EXPOSE 8080
CMD ["./wordbook-generator"]
