FROM golang:1.18

WORKDIR /usr/src/app

COPY go.mod go.sum ./

COPY ./.docker/golang/entrypoint.sh /usr/scr/app/.docker/golang/entrypoint.sh

RUN go mod download && go mod verify

RUN curl -s https://packagecloud.io/install/repositories/golang-migrate/migrate/script.deb.sh | bash
RUN apt-get update && apt-get install -y migrate 
RUN chmod +x /usr/scr/app/.docker/golang/entrypoint.sh