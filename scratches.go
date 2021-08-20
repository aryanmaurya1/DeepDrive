package main

import (
	"log"
)

func checkError(err error) {
	if err != nil {
		log.Println(err)
	}
}

// TASK 1 : Reading an image, and then writing it into a file.
// func main() {
// 	img, err := os.Open("assets/1.png")
// 	checkError(err)
// 	fileInfo, err := img.Stat()
// 	checkError(err)
// 	var size = fileInfo.Size()
// 	var buffer []byte = make([]byte, size)
// 	img.Read(buffer)
// 	img.Close()

// 	outFile, err := os.Create("assets/img.png")
// 	checkError(err)

// 	n, err := outFile.Write(buffer)
// 	outFile.Close()
// 	checkError(err)
// 	fmt.Println(n)
// }

// TASK 2 : BREAKING AN IMAGE INTO MULTIPLE FILES 
// func main() {

// 	img, err := os.Open("assets/1.png")
// 	checkError(err)
// 	defer img.Close()

// 	fileInfo, err := img.Stat()
// 	checkError(err)

// 	size := fileInfo.Size()
// 	buffer := make([]byte, size)

// 	img.Read(buffer)
// 	os.Mkdir("assets/img_1", os.FileMode(0777))
// 	var singleFileSize int64 = size / 10

// 	count := int64(0)
// 	for i := int64(0); i < size; {
// 		var singleChunk []byte
// 		if (i + singleFileSize) < size {
// 			singleChunk = buffer[i : i+singleFileSize]
// 		} else {
// 			singleChunk = buffer[i:]
// 		}
// 		fmt.Println("Written range : ", i, i+singleFileSize)
// 		file, err := os.Create(fmt.Sprintf("assets/img_1/img_%d", count))
// 		checkError(err)
// 		file.Write(singleChunk)
// 		file.Close()
// 		i = i + singleFileSize
// 		count++
// 	}
// }

// TASK 3 : COMBINING IMAGE 
// func main() {
// 	files, err := os.ReadDir("assets/img_1")
// 	var imgName = "constructed_img.png"
// 	var finalBuffer []byte
// 	checkError(err)
// 	for index, file := range files {
// 		fmt.Println(file, index)

// 		singleChunk, err := os.Open(fmt.Sprintf("assets/img_1/%s", file.Name()))
// 		checkError(err)

// 		singleChunkSize, err := singleChunk.Stat()
// 		checkError(err)
// 		chunkLength := singleChunkSize.Size()
// 		tempBuffer := make([]byte, chunkLength)
// 		singleChunk.Read(tempBuffer)
// 		finalBuffer = append(finalBuffer, tempBuffer...)
// 		singleChunk.Close()
// 	}
// 	outFile, err := os.Create(imgName)
// 	checkError(err)
// 	defer outFile.Close()
// 	outFile.Write(finalBuffer)
// }


// TASK 4 : COMPARING IMAGE BYTES
// func main() {

// 	originalImg, err := os.Open("assets/1.png")
// 	checkError(err)
// 	defer originalImg.Close()

// 	info, err := originalImg.Stat()
// 	checkError(err)
	
// 	size := info.Size()
// 	originalBuffer := make([]byte, size)
// 	originalImg.Read(originalBuffer)


// 	// PART TWO 
// 	cImage, err := os.Open("constructed_img.png")
// 	checkError(err)
// 	defer cImage.Close()

// 	info, err = cImage.Stat()
// 	checkError(err)
	
// 	size = info.Size()
// 	cBuffer := make([]byte, size)
// 	cImage.Read(cBuffer)

// 	fmt.Println(len(originalBuffer), len(cBuffer))

// 	for index, value := range originalBuffer {
// 		if value != cBuffer[index] {
// 			fmt.Println("value mismatch : ", value, cBuffer[index])
// 		}
// 	} 
// }