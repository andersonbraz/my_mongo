package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"./helper"
	"./models"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
)

func getRoot(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// we created Info array
	intro := []string{"Golang", "Example API"}
	json.NewEncoder(w).Encode(intro) // encode similar to serialize process.

}

func getCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// we created Info array
	intro := []string{"Golang", "Example API"}
	json.NewEncoder(w).Encode(intro) // encode similar to serialize process.

}

func getInfos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// we created Info array
	var infos []models.Info

	//Connection mongoDB with helper class
	collection := helper.ConnectDB()

	// bson.M{},  we passed empty filter. So we want to get all data.
	cur, err := collection.Find(context.TODO(), bson.M{})

	if err != nil {
		helper.GetError(err, w)
		return
	}

	// Close the cursor once finished
	/*A defer statement defers the execution of a function until the surrounding function returns.
	simply, run cur.Close() process but after cur.Next() finished.*/
	defer cur.Close(context.TODO())

	for cur.Next(context.TODO()) {

		// create a value into which the single document can be decoded
		var info models.Info
		// & character returns the memory address of the following variable.
		err := cur.Decode(&info) // decode similar to deserialize process.
		if err != nil {
			log.Fatal(err)
		}

		// add item our array
		infos = append(infos, info)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(infos) // encode similar to serialize process.
}

// var client *mongo.Client

func main() {
	//Init Router
	r := mux.NewRouter()

	r.HandleFunc("/", getRoot).Methods("GET")
	r.HandleFunc("/api", getRoot).Methods("GET")
	r.HandleFunc("/api/check", getCheck).Methods("GET")
	r.HandleFunc("/api/infos", getInfos).Methods("GET")

	log.Fatal(http.ListenAndServe(":8030", r))

}
