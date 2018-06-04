package main

import (
	"encoding/json"
	"log"
	"net/http"
)

var state = map[string]bool{
	"1": true,
	"2": false,
	"3": true,
	"4": false,
}

func myHandler(w http.ResponseWriter, r *http.Request) {
	data, _ := json.Marshal(state)
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func main() {
	http.HandleFunc("/enable", myHandler)
	http.HandleFunc("/disable", myHandler)
	http.HandleFunc("/state", myHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
