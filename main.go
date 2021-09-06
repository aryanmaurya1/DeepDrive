package main

import (
	"fmt"
	"log"
	"net/http"
	"ourtool/internal/db"
	"ourtool/internal/server"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println(err.Error())
	}

	dbConn, err := db.GetConnection()
	if err != nil {
		log.Println("Connection to db failed !!")
	}
	dbConn.AutoMigrate(&db.User{}, &db.File{}, &db.Chunk{})

	r := server.InitRoutes()
	r.Use(server.Logger) // Attaching the middleware
	http.ListenAndServe("0.0.0.0:8080", r)

}
