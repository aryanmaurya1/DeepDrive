package main

import (
	"net/http"
	"ourtool/internal/server"
)

func main() {
	r := server.InitRoutes()
	http.ListenAndServe(":8080", r)
}
