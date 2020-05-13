package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	hosts      = "localhost:27017"
	database   = "dbinfo"
	username   = "appuser"
	password   = "M"
	collection = "info"
)

type Info struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type MongoStore struct {
	session *mgo.Session
}

var mongoStore = MongoStore{}

func main() {

	//Create MongoDB session
	session := initialiseMongo()
	mongoStore.session = session

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/infos", jobsGetHandler).Methods("GET")
	router.HandleFunc("/infos", jobsPostHandler).Methods("POST")

	log.Fatal(http.ListenAndServe(":8030", router))

}

func initialiseMongo() (session *mgo.Session) {

	info := &mgo.DialInfo{
		Addrs:    []string{hosts},
		Timeout:  60 * time.Second,
		Database: database,
		Username: username,
		Password: password,
	}

	session, err := mgo.DialWithInfo(info)
	if err != nil {
		panic(err)
	}

	return

}

func jobsGetHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")

	col := mongoStore.session.DB(database).C(collection)

	results := []Info{}
	col.Find(bson.M{"title": bson.RegEx{"", ""}}).All(&results)
	jsonString, err := json.Marshal(results)
	if err != nil {
		panic(err)
	}
	fmt.Fprint(w, string(jsonString))

}

func jobsPostHandler(w http.ResponseWriter, r *http.Request) {

	col := mongoStore.session.DB(database).C(collection)

	//Retrieve body from http request
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		panic(err)
	}

	//Save data into Info struct
	var _info Info
	err = json.Unmarshal(b, &_info)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	//Insert info into MongoDB
	err = col.Insert(_info)
	if err != nil {
		panic(err)
	}

	//Convert info struct into json
	jsonString, err := json.Marshal(_info)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	//Set content-type http header
	w.Header().Set("content-type", "application/json")

	//Send back data as response
	w.Write(jsonString)

}
