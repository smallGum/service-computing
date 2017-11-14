package main

import (
	"log"
	"os"

	"service-computing/cloudgo/server"

	"github.com/spf13/pflag"
)

// define default port
const DEFAULT_PORT string = "23333"

func main() {
	// get port from the environment variable
	// if not, set port to DEFAULT_PORT
	port := os.Getenv("DEFAULT_PORT")
	if len(port) == 0 {
		port = DEFAULT_PORT
	}

	// get port from parameter
	pPort := pflag.StringP("port", "p", DEFAULT_PORT, "PORT for http server listening")
	pflag.Parse()
	if len(*pPort) != 0 {
		port = *pPort
	}

	err := server.Serving(port)
	if err != nil {
		log.Fatal(err)
	}
}
