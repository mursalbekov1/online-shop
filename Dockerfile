FROM golang:1.22.5-alpine as builder

WORKDIR /app

RUN go mod download

COPY . .

RUN go build -o user-service .

FROM alpine as runner

COPY --from=builder /app/user-service .
COPY config/config.yaml ./config/config.yaml

CMD ["./user-service"]