FROM golang:alpine

RUN apk add --no-cache git

WORKDIR /app

COPY go.mod .

COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o main .

CMD ["/app/main"]
