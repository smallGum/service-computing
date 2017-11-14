package server

import (
	"net/http"
	"service-computing/cloudgo/server/mux"
)

// Serving create a server on the specific addr and listening
func Serving(addr string) error {
	server := &http.Server{
		Addr:    ":" + addr,
		Handler: mux.GetMux(),
	}

	return server.ListenAndServe()
}
