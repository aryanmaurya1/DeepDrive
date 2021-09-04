package main

import (
	"os"
	"ourtool/internal/core"
	"ourtool/internal/divmer"
	"path"
)

const filename string = "sample_executable.AppImage"
const readBaseDir string = "files/original"
const mergeBaseDir string = "files/merged"
const brokenBaseDir string = "files/broken"

func main() {
	logfile, _ := os.Create("logs")
	defer logfile.Close()
	img, err := core.ReadFile(path.Join(readBaseDir, filename))
	if err != nil {
		return
	}
	var writingConfig divmer.DivideAndWriteConfig

	writingConfig.BaseDirectory = brokenBaseDir
	writingConfig.FileName = filename
	writingConfig.Data = img
	writingConfig.BufferSize = 1024 * 1024 * 20 // size
	writingConfig.Metadata = nil
	writingConfig.OutputRef = logfile

	c := divmer.DivideAndWrite(writingConfig)
	names := core.GetOriginalNameList(c)

	var mergingConfig divmer.ReadAndMergeConfig

	mergingConfig.BaseDirectory = brokenBaseDir
	mergingConfig.WritePath = mergeBaseDir
	mergingConfig.FileName = filename
	mergingConfig.ReadingOrder = names

	divmer.ReadAndMerge(mergingConfig)
}
