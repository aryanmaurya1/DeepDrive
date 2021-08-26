package main

import (
	"ourtool/internal/core"
	"ourtool/internal/divmer"
)

func main() {
	img := core.ReadFile("assets/test.png")
	var writingConfig divmer.DivideAndWriteConfig
	
	writingConfig.BaseDirectory = "assets"
	writingConfig.FileName = "new.png"
	writingConfig.Data = img
	writingConfig.BufferSize = 1024 * 1024 * 5 // 5MB
	writingConfig.Metadata = nil

	divmer.DivideAndWrite(writingConfig)
}