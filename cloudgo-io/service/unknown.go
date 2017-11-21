package service

import (
	"fmt"
	"net/http"
)

// Error replies to the request with the specified error message and HTTP code.
// It does not otherwise end the request; the caller should ensure no further
// writes are done to w.
// The error message should be plain text.
func Error(w http.ResponseWriter, error string, code int) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(code)
	fmt.Fprintln(w, error)
}

// NotImplemented replies to the request with an HTTP 501 not Implemented error.
func NotImplemented(w http.ResponseWriter, r *http.Request) {
	Error(w, "501 Not Implemented", http.StatusNotImplemented)
}
