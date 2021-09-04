package server

import (
	"log"
	"os"
)

type CustomError struct {
	Result string `json:"result"`
	Msg    string `json:"msg"`
}

type SuccessResponse struct {
	Result string `json:"result"`
	Msg    string `json:"msg"`
}

// SetOutputRef : Setting custom logging files.
func SetOutputRef(o *os.File) {
	OutputRef = o
}

func checkError(err error) *CustomError {
	var customErr *CustomError = nil
	if err != nil {
		log.Println(err.Error())
		customErr = &CustomError{Result: "fail", Msg: err.Error()}
	}
	return customErr
}
