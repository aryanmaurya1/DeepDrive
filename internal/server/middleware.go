package server

import (
	"log"
	"net/http"
)

// Logger : Currently this is a simple logger,
// later we will increase it logging capabilities.
func Logger(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Method, r.URL)
		h.ServeHTTP(w, r)
	})
}
