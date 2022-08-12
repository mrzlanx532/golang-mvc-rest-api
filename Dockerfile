FROM golang:1.18

WORKDIR /usr/src/app

COPY go.mod go.sum ./

RUN go mod download && go mod verify && apt-get update

RUN curl -s https://packagecloud.io/install/repositories/golang-migrate/migrate/script.deb.sh | bash
RUN apt-get update && apt-get install -y migrate 