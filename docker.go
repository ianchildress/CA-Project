package main

import (
	"github.com/fsouza/go-dockerclient"
	"time"
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

func listImages() (images []docker.APIImages, err error) {
	client, _ := docker.NewClient(SOCKET)
	return client.ListImages(docker.ListImagesOptions{All: false})
}

func listContainers() (containers []docker.APIContainers, err error) {
	client, _ := docker.NewClient(SOCKET)
	return client.ListContainers(docker.ListContainersOptions{All: true})
}
