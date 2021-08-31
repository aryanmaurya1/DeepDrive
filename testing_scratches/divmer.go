package main

import (
	"fmt"
	"os"
	"ourtool/internal/core"
	"ourtool/internal/divmer"
)

func main() {
	logfile, _ := os.Create("logs")
	defer logfile.Close()
	img := core.ReadFile("assets/2.png")
	var writingConfig divmer.DivideAndWriteConfig

	writingConfig.BaseDirectory = "assets"
	writingConfig.FileName = "2.png"
	writingConfig.Data = img
	writingConfig.BufferSize = 1024 * 256 // 5MB
	writingConfig.Metadata = nil
	writingConfig.OutputRef = logfile

	c := divmer.DivideAndWrite(writingConfig)
	for index, value := range c {
		fmt.Println(index, value)
	}
}

// func main() {
// 	logfile, _ := os.Create("logs")
// 	defer logfile.Close()
// 	var mergingConfig divmer.ReadAndMergeConfig

// 	mergingConfig.BaseDirectory = "assets"
// 	mergingConfig.FileName = "2.png"
// 	mergingConfig.OutputRef = logfile
// 	mergingConfig.MetadataSize = 0

// 	divmer.ReadAndMerge(mergingConfig)
// }
