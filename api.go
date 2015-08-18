package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/fsouza/go-dockerclient"
	"github.com/julienschmidt/httprouter"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

func serve() {
	router := httprouter.New()

	// GET
	router.GET("/containers", apiGetContainers)
	router.GET("/images", apiGetImages)
	router.GET("/container/:id/start", apiStartContainer)
	router.GET("/container/:id/stop", apiStopContainer)
	router.GET("/container/:id/inspect", apiInspectContainer)
	router.ServeFiles("/static/*filepath", http.Dir("static/app"))

	// POST

	log.Fatal(http.ListenAndServe(":8080", router))
}

func apiStartContainer(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	err := startContainer(id, &docker.HostConfig{})

	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(ErrorResponse{Error: err.Error()})
		log.Printf("%v failed to start.", id)

	} else {
		w.WriteHeader(200)
		log.Printf("%v started.", id)
	}
}

func apiInspectContainer(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	container, err := inspectContainer(id)

	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(ErrorResponse{Error: err.Error()})
		log.Printf("%v failed to inspect. Error: %v", id, err)

	} else {
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(container)
		log.Printf("%v inspected.", id)
	}
}

func apiStopContainer(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	err := stopContainer(id)

	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(ErrorResponse{Error: err.Error()})
		log.Printf("%v failed to stop.", id)
	} else {
		w.WriteHeader(200)
		log.Printf("%v stopped.", id)
	}
}

func apiGetContainers(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Connect to the DOCKER HOST via unix socket and get a list of our containers.
	// We'll put them into our database.
	containers, err := listContainers(docker.ListContainersOptions{All: true})
	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(ErrorResponse{Error: err.Error()})
	} else {
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(containers)
	}
}

func apiGetImages(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Connect to the DOCKER HOST via unix socket and get a list of our containers.
	// We'll put them into our database.
	images, err := listImages(docker.ListImagesOptions{All: false})
	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(ErrorResponse{Error: err.Error()})
	} else {
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(images)
	}
}
