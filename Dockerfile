FROM golang:alpine AS builder

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o api ./cmd/api

FROM alpine

COPY --from=builder /app/api /app/api

EXPOSE 8080

ENTRYPOINT ["/app/api"]