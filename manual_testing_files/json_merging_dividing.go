package main

type FileChunk struct {
	Index int    `json:"index"`
	Data  []byte `json:"data"`
}

// func main() {
// 	var basePath = "assets"
// 	var fileName = "2.png"
// 	var directory = path.Join(basePath, fileName+"_dir")
// 	files, err := os.ReadDir(directory)
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// 	err = os.Mkdir(path.Join(basePath, fileName+"_jsondir"), 0777)
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// 	for index, value := range files {
// 		var chunkReadPath = path.Join(directory, value.Name())
// 		var chunkWritePath = path.Join(basePath, fileName+"_jsondir", value.Name())
// 		data, err := os.ReadFile(chunkReadPath)
// 		if err != nil {
// 			log.Fatalln(err)
// 		}
// 		c := FileChunk{Data: data, Index: index}
// 		jsonChunk, err := json.Marshal(c)
// 		if err != nil {
// 			log.Fatalln(err)
// 		}
// 		writeFile, err := os.Create(chunkWritePath)
// 		if err != nil {
// 			log.Fatalln(err)
// 		}
// 		writeFile.Write(jsonChunk)
// 		writeFile.Close()
// 		fmt.Println(index,chunkReadPath, chunkWritePath)
// 	}
// }

// func main() {
// 	var basePath = "assets"
// 	var fileName = "2.png"
// 	files, err := os.ReadDir(path.Join(basePath, fileName+"_jsondir"))
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// 	var fullFile []byte
// 	for _, chunk := range files {
// 		var chunkFileReadPath = path.Join(basePath, fileName+"_dir", chunk.Name())
// 		chunk, err := os.ReadFile(chunkFileReadPath)
// 		if err != nil {
// 			log.Fatalln(err)
// 		}
// 		var singleChunk FileChunk
// 		err = json.Unmarshal(chunk, &singleChunk)
// 		if err != nil {
// 			log.Fatalln(err)
// 		}
// 		// fmt.Println(singleChunk.Data, chunk)
// 		fullFile = append(fullFile, singleChunk.Data...)
// 	}
// 	mergedFile, err := os.Create(path.Join(basePath, "merged_"+fileName))
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// 	fmt.Println(fullFile)
// 	mergedFile.Write(fullFile)
// 	mergedFile.Close()
// }
