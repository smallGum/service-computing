package service

import (
	"log"
	"net/http"

	"github.com/gorilla/schema"
	"github.com/unrolled/render"
)

// Student student structure
type Student struct {
	ID       string `json:"id"`
	UserName string `json:"username"`
	Subject  string `json:"subject"`
	Score    string `json:"score"`
}

// deal with data from form
func formHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		student := new(Student)
		decoder := schema.NewDecoder()
		// req.PostForm is a map of our POST form values
		err := decoder.Decode(student, req.PostForm)
		if err != nil {
			log.Fatal(err)
		}

		formatter.HTML(w, http.StatusOK, "table", *student)
	}
}
