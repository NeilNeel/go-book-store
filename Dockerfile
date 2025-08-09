FROM golang:1.23-alpine

WORKDIR /app

COPY go.mod .
COPY go.sum .

COPY . .

RUN go build -o book-server .

EXPOSE 3000

CMD ["./book-server"]
