FROM golang:latest

WORKDIR /go/src/github.com/leogaonabr/golang-service
ENV ENV=$ENV
ENV PORT=$PORT

COPY . .

RUN go build -o main .

EXPOSE 9000
ENTRYPOINT ["./main", "server"]