## go-hello-server

> A simple http server in go.

### Setup

* Install go `1.14` or higher from [here](https://golang.org/doc/install)
* GNU Make
* Execute `make build`

### Run

* Execute `make run CMD="server"`

##### Docker

Pre-requisite is to install Docker on your machine. [link](https://docs.docker.com/engine/install/)

* Build the image using `docker build -t go-hello-server .`
* Run container using `docker run -p 9000:9000 -d go-hello-server`
* Your server runs on a container and exposed via 9000 port

### Contributors

Harsha 
