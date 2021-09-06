package main

import (
	"fmt"
	"ourtool/internal/db"

	"github.com/joho/godotenv"
)

func main() {

	godotenv.Load(".env")
	dbconn, err := db.GetConnection()
	if err != nil {
		fmt.Println(err.Error())
	}
	var user db.User
	var file db.File

	user.Name = "Aryan"
	user.Username = "aryanmaurya1"
	user.Key = "12341234"
	user.Password = "abcdefgh"

	file.User = user
	file.OriginalName = "file.jsp"
	file.Hash = "hasher"
	file.Size = 9087
	file.ChunkCount = 10

	dbconn.AutoMigrate(&db.User{}, &db.File{}, &db.Chunk{})
	err = dbconn.Debug().Model(&db.File{}).Create(&file).Error

}
