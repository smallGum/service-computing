package service

import (
	"net/http"

	"github.com/unrolled/render"
)

// render the home page and write to the response
func homeHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		formatter.JSON(w, http.StatusOK, struct {
			Author string `json:"author"`
			OS     string `json:"os"`
			Date   string `json:"date"`
		}{})
	}
}
