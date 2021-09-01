package main

import (
	"fmt"
	"os"
	"ourtool/internal/core"
	"ourtool/internal/divmer"
	"path"
)

const filename string = "Geoffrey Lessel - Phoenix in Action (2018, Manning).pdf"
const readBaseDir string = "original"
const mergeBaseDir string = "merged"
const brokenBaseDir string = "broken"

func main() {
	logfile, _ := os.Create("logs")
	defer logfile.Close()
	img := core.ReadFile(path.Join(readBaseDir, filename))
	var writingConfig divmer.DivideAndWriteConfig

	writingConfig.BaseDirectory = brokenBaseDir
	writingConfig.FileName = filename
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

	mergingConfig.BaseDirectory = brokenBaseDir
	mergingConfig.WritePath = mergeBaseDir
	mergingConfig.FileName = filename
	mergingConfig.ReadingOrder = names

	divmer.ReadAndMerge(mergingConfig)
}
