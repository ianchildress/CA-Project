package main

import (
	"encoding/json"
	"github.com/fsouza/go-dockerclient"
	"io/ioutil"
	"log"
	"strconv"
	"time"
)

type ConfigFile struct {
	Start   StartConfig `json:"start"`
	Configs HostTypes   `json:"configs"`
}

// ========= Start section =====================
type StartConfig struct {
	Images     []string      `json:"images"`
	Containers []string      `json:"containers"`
	Settings   StartSettings `json:"settings"`
}

type StartSettings struct {
	StopContainers bool `json:"stop-containers"`
}

// ========= Configs section =====================
type HostTypes struct {
	Images     []ImageConfig      `json:"images"`
	Containers []ContainerConfigs `json:"containers"`
}

// ========= Images ==============================
type ImageConfig struct {
	Image   string                        `json:"image"`
	Options docker.CreateContainerOptions `json:"options"`
}

// ========= Containers ==========================
type ContainerConfigs struct {
	Id         string             `json:"id"`
	Hostconfig *docker.HostConfig `json:"hostconfig"`
}

func LoadConfigFile() (config ConfigFile) {
	f, err := ioutil.ReadFile("config.json")
	if err != nil {
		log.Panic(err)
	}

	err = json.Unmarshal(f, &config)
	if err != nil {
		log.Printf("Error loading config file. %v", err)
	}

	return
}

func Start(config ConfigFile) {
	// Start up parameters
	if config.Start.Settings.StopContainers {
		StopContainers()
	}
	// Load containers
	// We will iterate through the specified containers and look for a matching
	// docker.HostConfig if it is specified in the config.

	// Iterate over containers to start up
	for i := range config.Start.Containers {
		var hostConfig *docker.HostConfig
		for j := range config.Configs.Containers {
			// Iterate over specified HostConfigs and look for a match
			if config.Configs.Containers[j].Id == config.Start.Containers[i] {
				hostConfig = config.Configs.Containers[j].Hostconfig
				break
			}
		}
		// Start the container
		err := startContainer(config.Start.Containers[i], hostConfig)
		if err != nil {
			log.Printf("Error starting container %v. Error: %v", config.Configs.Containers[i].Id, err)
		} else {
			log.Printf("Container %v started successfully.", config.Configs.Containers[i].Id)
		}
	}

	for i := range config.Start.Images {
		var opts docker.CreateContainerOptions
		var hostConfig *docker.HostConfig

		for j := range config.Configs.Images {
			hostConfig = &docker.HostConfig{}
			// Iterate over specified HostConfigs and look for a match
			if config.Configs.Images[j].Image == config.Start.Images[i] {
				opts = config.Configs.Images[j].Options
				opts.Name = config.Configs.Images[j].Image + strconv.FormatInt(time.Now().UnixNano(), 10)
				hostConfig = config.Configs.Images[j].Options.HostConfig
				break
			}
		}
		// Create a container

		container, err := createContainer(opts)
		if err != nil {
			log.Panic(err)
		}

		err = startContainer(container.ID, hostConfig)

		if err != nil {
			log.Printf("Error creating Image %v. Error: %v", config.Configs.Images[i].Image, err)
		} else {
			log.Printf("Image %v started successfully as %v.", container.Image, container.Name)
		}
	}
}
