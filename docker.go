package main

import (
	"log"
	"time"

	"github.com/fsouza/go-dockerclient"
)

type container struct {
	Container_id string
	Image        string
	Command      string
	Created      time.Time
	Status       string
	Ports        string
	Names        string
}

func StartImage(containerType string) (id string, err error) {
	config := &docker.HostConfig{}

	switch containerType {
	case "web":
		config.Links = append(config.Links, "civis-mysql")
		id = "web1"
	case "mysql":
		id = "civis-mysql"
	}

	err = startContainer(id, config)
	return
}

func StopContainers() {
	containers, err := listContainers(docker.ListContainersOptions{All: false})
	if err != nil {
		log.Panic(err)
	}

	if len(containers) == 0 {
		return
	}

	// Iterate over the containers and stop them.
	for i := range containers {
		err = stopContainer(containers[i].ID)
		if err != nil {
			log.Panic(err)
		}
		log.Printf("%v stopped.", containers[i].Names)
	}

	return
}

func listImages(options docker.ListImagesOptions) (images []docker.APIImages, err error) {
	client, _ := docker.NewClient(SOCKET)
	return client.ListImages(options)
}

func listContainers(options docker.ListContainersOptions) (containers []docker.APIContainers, err error) {
	client, _ := docker.NewClient(SOCKET)
	return client.ListContainers(options)
}

func createContainer(opts docker.CreateContainerOptions) (*docker.Container, error) {
	client, _ := docker.NewClient(SOCKET)
	return client.CreateContainer(opts)
}

func startContainer(id string, config *docker.HostConfig) error {
	client, _ := docker.NewClient(SOCKET)
	return client.StartContainer(id, config)
}

func stopContainer(id string) error {
	client, _ := docker.NewClient(SOCKET)
	return client.StopContainer(id, 15)
}
