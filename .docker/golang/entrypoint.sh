#!/bin/bash

while true
do 
    ERROR=$(migrate -source file:./database/migrations -database mysql://golang_rest_api:golang_rest_api@tcp\(mysql:3306\)/golang_rest_api up 2>&1)
    echo $ERROR

    if [[ ! $ERROR =~ "connection refused" ]]; then
        break
    fi
    sleep 5
done

go run /usr/src/app/cmd/app/main.go