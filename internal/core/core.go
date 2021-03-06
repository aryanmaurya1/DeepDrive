package core

import (
	"log"
)

// THIS FILE CONTAINS GENERAL FUNCTIONS WHICH DO NOT PRODUCE SIDE EFFECTS

func checkError(err error) {
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func GetSizeFromLength(buffer []byte) (int, int) {
	size := len(buffer)
	// returns size in Kb and Mb
	return size / (1024), size / (1024 * 1024)
}

func GetOriginalNameList(c []FileChunk) []string {
	var names []string
	for _, name := range c {
		names = append(names, name.OriginalName)
	}
	return names
}

func GetEncNameList(c []FileChunk) []string {
	var names []string
	for _, name := range c {
		names = append(names, name.EncryptedName)
	}
	return names
}
