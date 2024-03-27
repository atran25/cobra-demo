FROM golang:1.22

WORKDIR /usr/src/app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY ./ ./
RUN go build -v -o /usr/local/bin/app ./

COPY --from=1password/op:2 /usr/local/bin/op /usr/local/bin/op

RUN op inject -i .env.tmpl -o .env

CMD ["app"]