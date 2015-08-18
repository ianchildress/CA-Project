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
	serve()

}
