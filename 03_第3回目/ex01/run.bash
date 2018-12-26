#!/bin/bash
go run ./main.go &

echo "curl http://localhost:8000"
curl http://localhost:8000

echo "curl http://localhost:8000/hello"
curl http://localhost:8000/hello

echo "curl http://localhost:8000/kon_san"
curl http://localhost:8000/kon_san

echo "curl http://localhost:8000/kon_san/majikayo"
curl http://localhost:8000/kon_san/majikayo

# ToDo:あとで止められるように修正
ps -ef|grep main.go |grep -v grep |xargs kill -9