package core

import (
	"log"
	"os"
)

// THIS FILE ONLY CONTAINS IO FUNCTIONS WHICH MAY PRODUCE SIDE EFFECTS

func ReadFile(filepath string) []byte {
	file, err := os.OpenFile(filepath, os.O_RDONLY, os.FileMode(1))
	checkError(err)
	defer file.Close()
	fileStat, err := file.Stat()
	checkError(err)
	fileSize := fileStat.Size()

	buffer := make([]byte, fileSize)
	n, err := file.Read(buffer)
	checkError(err)
	if n != int(fileSize) {
		log.Fatalf("Error in reading complete data, total len : %d Read len : %d", fileSize, n)
	}
	return buffer
}

func WriteToFile(data []byte, filepath string) *os.File {
	file, err := os.Create(filepath)
	checkError(err)

	n, err := file.Write(data)
	checkError(err)

	if n != len(data) {
		log.Fatalf("Error in writing complete data, total len : %d Written len : %d", len(data), n)
	}
	return file
}
