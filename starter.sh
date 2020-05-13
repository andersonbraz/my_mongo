#!/bin/zsh

## TEST 01
# docker-compose -f "docker-compose.yml" up -d --build
# go run dump.go
# docker build -t mongo-client:1.0 . -f Dockerfile
# docker run -d --name mongo-client -p 8030:8030 mongo-client:1.0

## TEST 02
docker build -t mongo-client:1.0 . -f Dockerfile
docker-compose -f "docker-compose.yml" up -d --build
go run dump.go