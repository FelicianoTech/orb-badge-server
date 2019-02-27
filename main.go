package main

import (
	"os"
	"strconv"

	log "github.com/sirupsen/logrus"
)

func main() {

	var port string

	if len(os.Args) == 2 {
		port = os.Args[1]
	} else {
		port = "1107"
	}

	if portInt, err := strconv.ParseInt(port, 10, 32); err != nil {
		log.Fatal("Invalid port number provided: " + port)
	} else {

		if portInt < -1 || portInt > 65535 {
			log.Fatal("Invalid port number provided: " + port)
		}
	}

	log.Info("Starting server...")

	server := Server{}
	server.initialize(false)

	log.Info("Server initialized.")

	server.run(":" + port)
}
