package server

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"ourtool/internal/core"
	"path"

	"github.com/gorilla/mux"
)

const FILE_UPLOAD_PATH = "original"

var OutputRef = os.Stdout

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	response := `{"msg":"Welcome to OurTool"}`
	fmt.Fprint(w, response)
}

func UploadFile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	userid := r.FormValue("userid")
	file, header, err := r.FormFile("file")
	checkError(err)

	fmt.Fprintf(OutputRef, "%s \n", userid) // Currently Just writing it to OutputRef
	fileData, err := ioutil.ReadAll(file)
	checkError(err)

	actualFileName := header.Filename
	go core.WriteToFile(fileData, path.Join(FILE_UPLOAD_PATH, actualFileName))

	response := `{msg:"file upload success", "result":"success"}`
	fmt.Fprint(w, response)
}

func InitRoutes() *mux.Router {
	var r = mux.NewRouter()
	r.HandleFunc("/", IndexHandler).Methods("GET")
	r.HandleFunc("/upload", UploadFile).Methods("POST")
	return r
}
