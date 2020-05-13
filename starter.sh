#!/bin/zsh
# Step 1 - Create Image
docker build -t mongo-client:1.0 . -f Dockerfile
# Step 2 - Up Containers
docker-compose -f "docker-compose.yml" up -d --build
# Step 3 - Populate Database
go run dump.go