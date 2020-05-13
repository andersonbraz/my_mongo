package main

import (
	"context"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Info struct {
	Title       string
	Description string
}

func main() {
	http.HandleFunc("/", infos)
	http.ListenAndServe(":8030", nil)
}

func infos(w http.ResponseWriter, r *http.Request) {

	clientOptions := options.Client().ApplyURI("mongodb://mongodb:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)

	collection := client.Database("dbinfo").Collection("infos")
	input := Info{"DevOps", "Example Service Golang with MongoDB"}
	rec := collection.InsertOne(context.TODO(), input)

	if err != nil {
		errin := []byte(`{ "infos": { "message": $err } }`)
		w.Header().Set("Content-Type", "application/json")
		w.Write(errin)
	}

	in := []byte(`{ "infos": { "message": "Connected." } }`)
	w.Header().Set("Content-Type", "application/json")
	w.Write(in)

}
