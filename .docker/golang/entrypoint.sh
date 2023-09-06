#!/bin/bash

while true
do 
    ANY_ANSWER=$(migrate -source file:./database/migrations -database mysql://golang_rest_api:golang_rest_api@tcp\(mysql:3306\)/golang_rest_api up 2>&1)
    echo "$ANY_ANSWER"

    if [[ ! $ANY_ANSWER =~ "connection refused" ]]; then
        break
    fi
    sleep 5
done

go build -o /usr/src/app/cmd/api/main /usr/src/app/cmd/api/main.go

/usr/src/app/cmd/api/main