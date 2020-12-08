#!/bin/bash

for i in {01..24}
do
    FILE=day_$i/main.go
    if test -f "$FILE"; then
        echo "Day $i"
        go run day_$i/main.go
        printf "\n"
    fi
done
