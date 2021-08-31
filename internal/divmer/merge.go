package divmer

import (
	"fmt"
	"os"
	"path"
	"sort"
)

type ReadAndMergeConfig struct {
	MetadataSize  int                 `json:"metadata"`
	WritePath     string              `json:"writePath"`
	BaseDirectory string              `json:"baseDirectory"`
	FileName      string              `json:"filename"`
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
	folderPath := path.Join(config.BaseDirectory, (config.FileName + "_dir"))
	files, err := os.ReadDir(folderPath)
	checkError(err)

	sort.SliceStable(files, func(i, j int) bool {
		i_info, _ := files[i].Info()
		j_info, _ := files[j].Info()
		return i_info.ModTime().UnixNano() < j_info.ModTime().UnixNano()
	})
	for i, v := range files {
		info, _ := v.Info()
		fmt.Println(i, info.ModTime().UnixNano(), v.Name())
	}
	return nil
}
