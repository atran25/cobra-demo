FROM golang:1.22

WORKDIR /app

COPY . .
RUN go build -v -o ./bin/api ./cmd/api

CMD ["/app/bin/api", "start"]