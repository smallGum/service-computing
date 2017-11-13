package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

/**
* sayHello is a handler function used to parse and response http request
* @param w used to write response to http client
* @param r the http request the server receives from the client
 */
func sayHello(w http.ResponseWriter, r *http.Request) {
	// parse the query parameters of the request and store the parameters in r.Form
	r.ParseForm()
	// print http request information on the server screen
	fmt.Println("path: ", r.URL.Path)
	fmt.Println("scheme: ", r.URL.Scheme)
	for k, v := range r.Form {
		fmt.Println("query parameter: ")
		fmt.Println(k+":", strings.Join(v, ""))
	}

	// write response to client
	fmt.Fprintf(w, "Hello, my client!\n")
}

func main() {
	// set the path of the server's address and set the handler function on the path
	http.HandleFunc("/", sayHello)
	// set the listening port
	err := http.ListenAndServe(":23333", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
