FROM golang:1.20.1-alpine3.17 AS builder
WORKDIR /app
COPY . .
RUN go build -o app .

FROM alpine:3.17 AS final
WORKDIR /app
COPY --from=builder /app/app .
CMD ["bash", "-c", "./app/app"]
