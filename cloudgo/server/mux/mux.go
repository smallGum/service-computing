package mux

import (
	"fmt"
	"net/http"
)

/**
* sayHello is a handler function used to say hello when user didn't tell his name
* @param w used to write response to http client
* @param r the http request the server receives from the client
 */
func sayHello(w http.ResponseWriter, r *http.Request) {
	// say hello to client
	if name := r.URL.Query().Get("name"); len(name) == 0 {
		fmt.Fprintf(w, "Hello, my client!\n")
	} else {
		fmt.Fprintf(w, "Hello, %s!\n", name)
	}
}

// GetMux return the router
func GetMux() *http.ServeMux {
	mx := http.NewServeMux()

	// set router's handlers
	mx.HandleFunc("/hello", sayHello)
	mx.HandleFunc("/hello/", sayHello)
	return mx
}
