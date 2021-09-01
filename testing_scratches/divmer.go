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
	names := core.GetOriginalNameList(c)

	var mergingConfig divmer.ReadAndMergeConfig

	mergingConfig.BaseDirectory = "assets"
	mergingConfig.WritePath = "assets/merging_dir"
	mergingConfig.FileName = "2.png"
	mergingConfig.ReadingOrder = names

	divmer.ReadAndMerge(mergingConfig)
}
