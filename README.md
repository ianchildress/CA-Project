# Polydocker
Polymer based front end for Docker. Manage Docker containers and images using a web interface, served by Go.

# Summary
This project uses a combination of Go, Docker, Polymer, and Javascript to create a simple interface for interacting with Docker hosts. 

## Start up
The app has a config.json file that allows the following:

* Start containers at start up
* Create containers from images at start up
* Stop and/or delete containers at start up

The app has a web interface that allows the following:

* Start/Stop containers
* View status of containers
* Create container from image

## Coming features
* Multiple host management
* Search and pull images from Docker Hub
* Inspect and details on containers
