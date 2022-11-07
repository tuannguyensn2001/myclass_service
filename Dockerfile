FROM golang:1.18-alpine3.16 as builder
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
COPY . .
RUN go build -o main src/server/main.go

FROM alpine:3.16
WORKDIR /app
COPY --from=builder /app/main .
COPY --from=builder /app/config.yml .

CMD ["/app/main","server"]