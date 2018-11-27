#!/bin/bash

echo "DEFAULT"
go run ./main.go
echo " "

echo "-flag1 200 を指定"
go run ./main.go -flag1 200
echo " "

echo "-flag2 昆さんまじかよ を指定"
go run ./main.go -flag2 昆さんまじかよ
echo " "

echo "-flag1 1000 を指定 -flag2 昆さん頼むぜ を指定"
go run ./main.go -flag2 昆さんまじかよ