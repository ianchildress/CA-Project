package main

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

func serve() {
	router := httprouter.New()
	router.GET("/containers", GetContainers)
	router.GET("/containers", GetImages)

	log.Fatal(http.ListenAndServe(":8080", router))
}

func GetContainers(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Connect to the DOCKER HOST via unix socket and get a list of our containers.
	// We'll put them into our database.
	containers, err := listContainers()
	if err != nil {
		w.WriteHeader(404)
		json.NewEncoder(w).Encode(ErrorResponse{Error: err.Error()})
	} else {
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(containers)
	}
}

func GetImages(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Connect to the DOCKER HOST via unix socket and get a list of our containers.
	// We'll put them into our database.
	images, err := listImages()
	if err != nil {
		w.WriteHeader(404)
		json.NewEncoder(w).Encode(ErrorResponse{Error: err.Error()})
	} else {
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(images)
	}
}
