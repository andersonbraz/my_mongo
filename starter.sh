#!/bin/zsh
## TEST 01
#docker-compose -f "docker-compose.yml" up -d --build
#docker build -t my-mongo:1.0 . -f Dockerfile
#docker run -d -p 8030:8030 --name mongo-service my-mongo:1.0 

## TEST 02
docker build -t mongodb-client:1.0 . -f Dockerfile
docker-compose -f "docker-compose-inc.yml" up -d --build