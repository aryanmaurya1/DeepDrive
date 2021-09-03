package server

import (
	"encoding/json"
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
	if e := checkError(err, w); e != nil {
		json.NewEncoder(w).Encode(e)
		return
	}
	defer file.Close()

	fmt.Fprintf(OutputRef, "%s \n", userid) // Currently Just writing it to OutputRef
	fileData, err := ioutil.ReadAll(file)
	if e := checkError(err, w); e != nil {
		json.NewEncoder(w).Encode(e)
		return
	}

	actualFileName := header.Filename
	err = core.WriteToFile(fileData, path.Join(FILE_UPLOAD_PATH, actualFileName))
	if e := checkError(err, w); e != nil {
		json.NewEncoder(w).Encode(e)
		return
	}

	response := SuccessResponse{Result: "success", Msg: "File upload sucessfull"}
	json.NewEncoder(w).Encode(response)
}

func InitRoutes() *mux.Router {
	var r = mux.NewRouter()
	r.HandleFunc("/", IndexHandler).Methods("GET")
	r.HandleFunc("/api/upload", UploadFile).Methods("POST")
	return r
}
