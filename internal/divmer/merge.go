package divmer

import (
	"fmt"
	"os"
	"path"
)

type ReadAndMergeConfig struct {
	MetadataSize  int                 `json:"metadata"`
	WritePath     string              `json:"writePath"`
	BaseDirectory string              `json:"baseDirectory"`
	FileName      string              `json:"filename"`
	ReadingOrder  []string            `json:"readingOrder"` // Order in which read of files must be performed
	OutputRef     *os.File            `json:"-"`
	PipeFn        func([]byte) []byte `json:"-"`
}

func ReadAndMerge(config ReadAndMergeConfig) *os.File {

	// setting default values of configuration if not provided
	if config.OutputRef == nil {
		config.OutputRef = os.Stdout
	}
	if config.PipeFn == nil {
		config.PipeFn = func(b []byte) []byte {
			return b
		}
	}
	if len(config.WritePath) == 0 {
		config.WritePath = config.BaseDirectory
	}
	var combinedFile []byte // This slice will hold combined file and will be written to disk
	folderPath := path.Join(config.BaseDirectory, (config.FileName + "_dir"))
	writePath := path.Join(config.WritePath, config.FileName)

	for _, filename := range config.ReadingOrder {
		chunkData, _ := os.ReadFile(path.Join(folderPath, filename))
		chunkData = config.PipeFn(chunkData)
		combinedFile = append(combinedFile, chunkData...)

		fmt.Fprintf(config.OutputRef, "Processed File : %s \n", path.Join(folderPath, filename))
	}
	writeFile, err := os.Create(writePath)
	checkError(err)
	writeFile.Write(combinedFile)
	return writeFile
}
