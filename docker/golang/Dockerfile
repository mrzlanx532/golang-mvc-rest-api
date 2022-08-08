FROM golang:1.18

WORKDIR /var/www/html

RUN go mod download && go mod verify

CMD ["go run main.go"]