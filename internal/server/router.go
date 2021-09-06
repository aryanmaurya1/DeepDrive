package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"ourtool/internal/core"
	"ourtool/internal/db"
	"path"
	"time"

	"github.com/gorilla/mux"
)

const FILE_UPLOAD_PATH = "files/original"

var OutputRef = os.Stdout

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	response := `{"msg":"Welcome to OurTool"}`
	fmt.Fprint(w, response)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {

	var user db.User
	json.NewDecoder(r.Body).Decode(&user)

	err := ValidateUserDetails(user) // from utils
	if err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}

	// every user will get unique key
	hashedPassword, errr := Hash(user.Password)
	if errr != nil {
		json.NewEncoder(w).Encode(SuccessResponse{"fail", errr.Error()})
		return
	}

	user.Password = string(hashedPassword)
	user.Key = fmt.Sprintf("%d_%d", time.Now().UnixNano(), rand.Int())

	errr = db.DB_CONNECTION.Debug().Create(&user).Error
	if errr != nil {
		json.NewEncoder(w).Encode(CustomError{"fail", errr.Error()})
		return
	}
	json.NewEncoder(w).Encode(SuccessResponse{"success", "user created"})

}

func UploadFile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	userid := r.FormValue("userid")
	file, header, err := r.FormFile("file")
	if e := checkError(err); e != nil {
		json.NewEncoder(w).Encode(e)
		return
	}
	defer file.Close()

	fmt.Fprintf(OutputRef, "USERID : %s \n", userid) // Currently Just writing it to OutputRef
	fileData, err := ioutil.ReadAll(file)
	if e := checkError(err); e != nil {
		json.NewEncoder(w).Encode(e)
		return
	}

	actualFileName := header.Filename
	err = core.WriteToFile(fileData, path.Join(FILE_UPLOAD_PATH, actualFileName))
	if e := checkError(err); e != nil {
		json.NewEncoder(w).Encode(e)
		return
	}

	response := SuccessResponse{Result: "success", Msg: "File upload sucessfull"}
	json.NewEncoder(w).Encode(response)
}

func InitRoutes() *mux.Router {
	var r = mux.NewRouter()
	r.HandleFunc("/", IndexHandler).Methods("GET")
	r.HandleFunc("/api/user/create", CreateUser).Methods("POST")
	r.HandleFunc("/api/upload", UploadFile).Methods("POST")
	return r
}
