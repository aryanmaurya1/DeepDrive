package divmer

import (
	"fmt"
	"os"
	"path"
)

type ReadAndMergeConfig struct {
	MetadataSize  int                 `json:"metadataSize"`  // Size of metadata which is prepended in each file chunk
	WritePath     string              `json:"writePath"`     // Path to the folder in which combined file must be written
	BaseDirectory string              `json:"baseDirectory"` // Base directory inside which the folder lies, which contains the broken chunks
	FileName      string              `json:"filename"`      // Name of the file which will store merged chunks
	ReadingOrder  []string            `json:"readingOrder"`  // Order in which read of the chunks must be performed
	OutputRef     *os.File            `json:"-"`             // Output file reference for writing logs
	PipeFn        func([]byte) []byte `json:"-"`             // Function through which every chunk will be passed before merging into a big file
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
