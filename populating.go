package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

func main() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017").
		SetAuth(options.Credential{
			AuthSource: "dbinfo", Username: "appuser", Password: "Mko0Zaq1",
		})

	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	var contentArray []interface{}

	contentArray = append(contentArray, bson.M{"Title": "DevOps - Alpha", "Description": "#01 Example by Anderson Braz."})
	contentArray = append(contentArray, bson.M{"Title": "DevOps - Bravo", "Description": "#02 Example by Anderson Braz."})
	contentArray = append(contentArray, bson.M{"Title": "DevOps - Charlie", "Description": "#03 Example by Anderson Braz."})
	contentArray = append(contentArray, bson.M{"Title": "DevOps - Delta", "Description": "#04 Example by Anderson Braz."})
	contentArray = append(contentArray, bson.M{"Title": "DevOps - Echo", "Description": "#05 Example by Anderson Braz."})
	contentArray = append(contentArray, bson.M{"Title": "DevOps - Fhox", "Description": "#06 Example by Anderson Braz."})
	contentArray = append(contentArray, bson.M{"Title": "DevOps - Golf", "Description": "#07 Example by Anderson Braz."})
	contentArray = append(contentArray, bson.M{"Title": "DevOps - Hotel", "Description": "#08 Example by Anderson Braz."})
	contentArray = append(contentArray, bson.M{"Title": "DevOps - India", "Description": "#09 Example by Anderson Braz."})
	contentArray = append(contentArray, bson.M{"Title": "DevOps - James", "Description": "#10 Example by Anderson Braz."})

	collection := client.Database("dbinfo").Collection("infos")

	if res, err := collection.InsertMany(context.TODO(), contentArray); err == nil {
		log.Println("Inserted record", res.InsertedIDs)
	} else {
		log.Fatal("Error inserting record:", err)
	}

}
