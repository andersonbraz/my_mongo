# Contêiner MongoDB + API Golang

## Activity

Evolve your Golang service to expose a /info API to HTTP GET requests, where the JSON content that gets returned is obtained from a MongoDB entry.

## Step 1 - Create Image

```shell
docker build -t mongo-client:1.0 . -f Dockerfile
```

## Step 2 - Up Containers

```shell
docker-compose -f "docker-compose.yml" up -d --build
```

## Step 3 - Populate Database

```shell
go run dump.go
```

## References

&#8658; [Deploys MongoDB with customization scripts and container with Mongo client](https://github.com/fabianlee/docker-mongodb/blob/master/docker-compose.yml)

&#8658; [MongoDB Examples With Golang](https://blog.ruanbekker.com/blog/2019/04/17/mongodb-examples-with-golang/)

&#8658; [Quick Start: Golang & MongoDB - How to Create Documents](https://www.mongodb.com/blog/post/quick-start-golang--mongodb--how-to-create-documents)

&#8658; [My First Go Microservice using MongoDB and Docker Multi-Stage Builds](https://www.melvinvivas.com/my-first-go-microservice/)

## History Command

```shell
clear
docker container ls -a
docker ps
docker container stop mongodb-server
docker container stop mongo-client
docker container prune -f
docker network prune -f
docker volume prune -f
docker container ls -a
docker ps
```

```shell
clear
docker ps
docker container ls -a
docker container prune -f
docker network prune -f
docker volume prune -f
docker container ls -a
docker ps
```