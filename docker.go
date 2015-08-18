package main

import (
	"github.com/fsouza/go-dockerclient"
)

func listImages(options docker.ListImagesOptions) ([]docker.APIImages, error) {
	client, _ := docker.NewClient(SOCKET)
	return client.ListImages(options)
}

func listContainers(options docker.ListContainersOptions) ([]docker.APIContainers, error) {
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

func removeContainer(opts docker.RemoveContainerOptions) error {
	client, _ := docker.NewClient(SOCKET)
	return client.RemoveContainer(opts)
}

func inspectContainer(id string) (*docker.Container, error) {
	client, _ := docker.NewClient(SOCKET)
	return client.InspectContainer(id)
}
