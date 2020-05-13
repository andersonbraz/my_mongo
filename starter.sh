#!/bin/zsh
docker build -t mongolang-client:1.0 . -f Dockerfile
docker run -d -p 8030:8030 --name my-mongolang-client mongolang-client:1.0
docker-compose -f "docker-compose.yml" up -d --build