FROM golang:1.22-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o stresstester

FROM alpine:3.18
WORKDIR /app
COPY --from=builder /app/stresstester .
RUN apk --no-cache add ca-certificates
ENTRYPOINT ["./stresstester"]