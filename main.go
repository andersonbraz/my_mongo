package main

import (
	"encoding/json"
	"html"
	"net/http"
)

type Response struct {
	Message string
}

func main() {
	http.HandleFunc("/", hello)
	http.ListenAndServe(":8030", nil)
}

func hello(w http.ResponseWriter, r *http.Request) {
	resp := Response{html.EscapeString(r.URL.Path)}

	js, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
