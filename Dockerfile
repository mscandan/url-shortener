# syntax=docker/dockerfile:1

FROM golang:alpine

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o out/url-shortener main.go  

EXPOSE 8080

CMD ["out/url-shortener"]
