package main

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
)

var state = map[string]bool{
	"1": false,
	"2": false,
	"3": false,
	"4": false,
}

func getValue(u *url.URL) string {
	q := u.Query()
	return q.Get("value")
}

func enable(w http.ResponseWriter, r *http.Request) {
	value := getValue(r.URL)
	if value == "all" {
		for socket, _ := range state {
			state[socket] = true
		}
	} else {
		state[value] = true
	}
	data, _ := json.Marshal(state)
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func disable(w http.ResponseWriter, r *http.Request) {
	value := getValue(r.URL)
	if value == "all" {
		for socket, _ := range state {
			state[socket] = false
		}
	} else {
		state[value] = false
	}
	data, _ := json.Marshal(state)
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func getstate(w http.ResponseWriter, r *http.Request) {
	data, _ := json.Marshal(state)
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func main() {
	http.HandleFunc("/enable", enable)
	http.HandleFunc("/disable", disable)
	http.HandleFunc("/state", getstate)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
