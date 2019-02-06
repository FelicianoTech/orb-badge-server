package main

import (
	log "github.com/sirupsen/logrus"
)

func main() {

	log.Info("Starting server...")

	server := Server{}
	server.initialize(false)

	log.Info("Server initialized.")

	server.run(":1107")
	//server.run(":80")
}
