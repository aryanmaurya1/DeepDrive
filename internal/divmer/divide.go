package divmer

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"ourtool/internal/core"
	"path"
	"time"
)

// Datatype for configuring the 'DivideAndWrite' function's working
// created them inside a single datatype because didn't wanted to pass
// 5 arguments to the function.
type DivideAndWriteConfig struct {
	Data          []byte              `json:"data"`          // Complete data
	Metadata      []byte              `json:"metadata"`      // Metadata slice to be prepended before writing slice to file
	BufferSize    int                 `json:"bufferSize"`    // Size of single file
	BaseDirectory string              `json:"baseDirectory"` // Base directory inside which to create folder to store broken files
	FileName      string              `json:"filename"`      // starting prefix of all broken files
	OutputRef     *os.File            `json:"-"`             // Output file reference for writing logs
	PipeFn        func([]byte) []byte `json:"-"`             // Function through which ever buffer will be passed before writing to file
}

func checkError(err error) {
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func DivideAndWrite(config DivideAndWriteConfig) []core.FileChunk {

	var chunks []core.FileChunk // This slice will be returned from function
	folderPath := path.Join(config.BaseDirectory, (config.FileName + "_dir"))
	size := len(config.Data)
	singleFileSize := config.BufferSize

	// setting default values of configuration if not provided
	if config.Metadata == nil {
		config.Metadata = []byte{}
	}
	if config.OutputRef == nil {
		config.OutputRef = os.Stdout
	}
	if config.PipeFn == nil {
		config.PipeFn = func(b []byte) []byte {
			return b
		}
	}

	os.Mkdir(folderPath, os.FileMode(0777))

	var count int
	for i := 0; i < size; {
		var singleChunk []byte
		var chunk core.FileChunk
		if (i + singleFileSize) < size {
			singleChunk = config.Data[i : i+singleFileSize]
		} else {
			singleChunk = config.Data[i:]
		}

		singleFilename := fmt.Sprintf("%s_%d", config.FileName, count)
		chunk.OriginalName = singleFilename                                     // storing original name of chunk
		chunk.Index = count                                                     // Storing ChunkId, it will be equal to 'count' not 'i'
		chunk.ChunkId = fmt.Sprintf("%d_%d", rand.Int(), time.Now().UnixNano()) // Setting ChunkId equal to random number plus its creation time

		file, err := os.Create(path.Join(folderPath, singleFilename))
		singleChunk = append(config.Metadata, singleChunk...)
		singleChunk = config.PipeFn(singleChunk)
		checkError(err)
		file.Write(singleChunk)
		file.Close()

		// --- Writting logs ---
		fmt.Fprint(config.OutputRef, "FILE : ", path.Join(folderPath, singleFilename))
		fmt.Fprintf(config.OutputRef, " || WRITTEN RANGE : %10d %10d \n", i, i+singleFileSize)

		i = i + singleFileSize
		count++
		chunks = append(chunks, chunk)
	}
	return chunks
}
