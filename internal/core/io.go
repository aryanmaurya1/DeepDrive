package core

import (
	"fmt"
	"os"
)

// THIS FILE ONLY CONTAINS IO FUNCTIONS WHICH MAY PRODUCE SIDE EFFECTS

func ReadFile(filepath string) ([]byte, error) {
	file, err := os.OpenFile(filepath, os.O_RDONLY, os.FileMode(1))
	if err != nil {
		return nil, err
	}
	defer file.Close()

	fileStat, err := file.Stat()
	if err != nil {
		return nil, err
	}
	fileSize := fileStat.Size()

	buffer := make([]byte, fileSize)
	n, err := file.Read(buffer)
	if err != nil {
		return nil, err
	}

	if n != int(fileSize) {
		return nil, fmt.Errorf("error in reading complete data, total len : %d Read len : %d", fileSize, n)
	}
	return buffer, nil
}

func WriteToFile(data []byte, filepath string) error {
	file, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer file.Close()

	n, err := file.Write(data)
	if err != nil {
		return err
	}

	if n != len(data) {
		return fmt.Errorf("error in writing complete data, total len : %d Written len : %d", len(data), n)
	}

	return nil
}
