package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"ourtool/internal/db"
	"ourtool/internal/server"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println(err.Error())
	}

	rand.Seed(time.Now().UnixNano())

	dbConn, err := db.GetConnection()
	db.DB_CONNECTION = dbConn // GLOBAL DB CONNECTION OBJECT

	if err != nil {
		log.Println("Connection to db failed !!")
	}
	dbConn.AutoMigrate(&db.User{}, &db.File{}, &db.Chunk{})

	r := server.InitRoutes()
	r.Use(server.Logger) // Attaching the middleware
	r.Use(server.ContentTypeJson)
	http.ListenAndServe("0.0.0.0:8080", r)

}
