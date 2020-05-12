# Contêiner MongoDB + API Golang

## Activity

Evolve your Golang service to expose a /info API to HTTP GET requests, where the JSON content that gets returned is obtained from a MongoDB entry.


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

OR

```shell
go run main.go
```

## Referências

&#8658; [Deploys MongoDB with customization scripts and container with Mongo client](https://github.com/fabianlee/docker-mongodb/blob/master/docker-compose.yml)

&#8658; [MongoDB Examples With Golang](https://blog.ruanbekker.com/blog/2019/04/17/mongodb-examples-with-golang/)

&#8658; [Quick Start: Golang & MongoDB - How to Create Documents](https://www.mongodb.com/blog/post/quick-start-golang--mongodb--how-to-create-documents)