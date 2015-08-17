package main

import (
	//"log"

	"github.com/fsouza/go-dockerclient"
)

const (
	SOCKET = "unix:///var/run/docker.sock"
)

var (
	started []docker.APIContainers
)

func main() {

	Start(LoadConfigFile())

	/*
	   StopContainers()
	   	if id, err := StartImage("mysql"); err != nil {
	   		log.Println(err)
	   	} else {
	   		log.Printf("%v started.", id)
	   	}
	   	if id, err := StartImage("web"); err != nil {
	   		log.Println(err)
	   	} else {
	   		log.Printf("%v started.", id)
	   	}

	   	serve()
	*/
	/*
		endpoint := "unix:///var/run/docker.sock"
		client, _ := docker.NewClient(endpoint)
		imgs, _ := client.ListImages(docker.ListImagesOptions{All: false})
		for _, img := range imgs {
			fmt.Println("ID: ", img.ID)
			fmt.Println("RepoTags: ", img.RepoTags)
			fmt.Println("Created: ", img.Created)
			fmt.Println("Size: ", img.Size)
			fmt.Println("VirtualSize: ", img.VirtualSize)
			fmt.Println("ParentId: ", img.ParentID)
		}
	*/
}
