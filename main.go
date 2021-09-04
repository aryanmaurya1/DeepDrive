package main

import (
	"net/http"
	"ourtool/internal/server"
)

func main() {
	r := server.InitRoutes()
	r.Use(server.Logger) // Attaching the middleware
	http.ListenAndServe("0.0.0.0:8080", r)
}
