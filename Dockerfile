FROM golang:latest

WORKDIR /app

COPY . .

RUN go mod tidy
RUN go mod vendor

RUN go build -o ./rssagg.exe ./cmd/main.go

EXPOSE 8000

CMD ["./rssagg.exe"]
