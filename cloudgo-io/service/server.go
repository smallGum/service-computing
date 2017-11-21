package service

import (
	"log"
	"net/http"
	"os"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
)

// get the html file from template directory
func getHTML() *render.Render {
	return render.New(render.Options{
		Directory:  "templates",
		Extensions: []string{".html"},
		IndentJSON: true,
	})
}

// initial the router of server
func initRouter(r *mux.Router, formatter *render.Render) {
	// get the virtual web root directory
	webRoot := os.Getenv("WEBROOT")
	if len(webRoot) == 0 {
		if root, err := os.Getwd(); err != nil {
			log.Fatal("could not retrive working directory!")
		} else {
			webRoot = root
		}
	}

	r.HandleFunc("/", homeHandler(formatter)).Methods("GET")
	r.HandleFunc("/register", formHandler(formatter)).Methods("POST")
	r.HandleFunc("/unknown", NotImplemented)
	r.PathPrefix("/").Handler(http.FileServer(http.Dir(webRoot + "/assets/")))
}

// NewServer create a new server
func NewServer() *negroni.Negroni {
	server := negroni.Classic()
	router := mux.NewRouter()
	formatter := getHTML()

	initRouter(router, formatter)
	server.UseHandler(router)

	return server
}
