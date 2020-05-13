# Contêiner MongoDB + API Golang

## Activity

Evolve your Golang service to expose a /info API to HTTP GET requests, where the JSON content that gets returned is obtained from a MongoDB entry.

## File: docker-compose.yml

```docker
version: '3.7'
services:
  my-mongodb:
    image: mongo
    container_name: my-mongodb
    ports:
      - 27017:27017
    environment:
      - MONGO_INITDB_DATABASE=info
      - MONGO_INITDB_ROOT_USERNAME=admin
      - MONGO_INITDB_ROOT_PASSWORD=Mko0Zaq1
    volumes:
      - ./mongo-entrypoint:/docker-entrypoint-initdb.d
      - mongodb:/data/db
      - mongoconfig:/data/configdb
    networks:
      - mongo_net
      
volumes:
  mongodb:
  mongoconfig:

networks:
  mongo_net:
    driver: bridge
```

## Execution Build Docker Compose

Execute command line:

```shell
docker-compose -f "docker-compose.yml" up -d --build
```

## Code: Golang in file main.go

```go
package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Person struct {
	Name string
	Age  int
	City string
}

func main() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

}
```

## Testing Code

Execute command line:

```shell
go run hello.go
```

## References

&#8658; [Deploys MongoDB with customization scripts and container with Mongo client](https://github.com/fabianlee/docker-mongodb/blob/master/docker-compose.yml)

&#8658; [MongoDB Examples With Golang](https://blog.ruanbekker.com/blog/2019/04/17/mongodb-examples-with-golang/)

&#8658; [Quick Start: Golang & MongoDB - How to Create Documents](https://www.mongodb.com/blog/post/quick-start-golang--mongodb--how-to-create-documents)

&#8658; [My First Go Microservice using MongoDB and Docker Multi-Stage Builds](https://www.melvinvivas.com/my-first-go-microservice/)

## History

clear
docker container ls -a
docker ps
docker container stop mongodb-server
docker container stop mongo-client
docker container prune
docker network prune
docker volume prune
docker container ls -a
docker ps