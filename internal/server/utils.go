package server

import (
	"log"
	"os"
)

// SetOutputRef : Setting custom logging files.
func SetOutputRef(o *os.File) {
	OutputRef = o
}

func checkError(err error) {
	if err != nil {
		log.Fatalln(err.Error())
	}
}
