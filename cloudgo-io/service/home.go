package service

import (
	"net/http"

	"github.com/unrolled/render"
)

// render the home page and write to the response
func homeHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		formatter.HTML(w, http.StatusOK, "index", struct {
			Author string `json:"author"`
			OS     string `json:"os"`
			Date   string `json:"date"`
		}{Author: "author: Jack Cheng", OS: "operating system: linux", Date: "data: 2017-11-20"})
	}
}
