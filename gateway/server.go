// Copyright (c) Alex Ellis 2017. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for full license information.

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
)

// Get device info with health data
func Device(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
}

// Get device Health info
func Health(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
}

// Route a request to a skill and return the response
func Skill(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

}

func register() error {

}

func main() {
	router := httprouter.New()
	router.GET("/", Device)
	router.GET("/health", Health)
	router.POST("/skill/:name", Skill)

	err := register()
	if err != nil {
		log.Fatal("Failed to register device with diot platform, error: ", err)
	}

	log.Fatal(http.ListenAndServe(":8080", router))
}
