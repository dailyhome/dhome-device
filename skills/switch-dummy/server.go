package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
)

var switch_nos = 0
var pin_map = map[string]int{}
var state = map[string]bool{}

func getValue(u *url.URL) string {
	q := u.Query()
	return q.Get("value")
}

func enable(w http.ResponseWriter, r *http.Request) {
	value := getValue(r.URL)
	if value == "all" {
		for socket, _ := range state {
			state[socket] = true
			log.Printf("Socket %s -> Pin %d is enabled", socket, pin_map[socket])
		}
	} else {
		state[value] = true
		log.Printf("Socket %s -> Pin %d is enabled", value, pin_map[value])
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
			log.Printf("Socket %s -> Pin %d is disabled", socket, pin_map[socket])
		}
	} else {
		state[value] = false
		log.Printf("Socket %s -> Pin %d is disabled", value, pin_map[value])
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

func initialize() {
	switch_nos, serr := strconv.Atoi(os.Getenv("switch_nos"))
	if serr != nil {
		log.Fatal("switch_nos cant be initialized, error :" + serr.Error())
	}
	log.Printf("No of switch %d", switch_nos)

	var pinArray []int = make([]int, switch_nos)
	pinmap_str := os.Getenv("switch_pins")
	jerr := json.Unmarshal([]byte(pinmap_str), &pinArray)
	if jerr != nil {
		log.Fatal("pin_map cant be initialized, error :" + serr.Error())
	}
	log.Printf("Pin Map %v", pinArray)

	i := 0
	for i < switch_nos {
		pin := fmt.Sprintf("%d", i+1)
		pin_map[pin] = pinArray[i]
		state[pin] = false
		i++
	}
}

func main() {

	initialize()

	http.HandleFunc("/enable", enable)
	http.HandleFunc("/disable", disable)
	http.HandleFunc("/state", getstate)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
